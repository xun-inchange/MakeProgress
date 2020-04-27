package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

//服务端

func main() {
	//代表一个tcp终端地址
	var tcpAddr *net.TCPAddr
	//ResolveTCPAddr将addr作为TCP地址解析并返回
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:999")

	//这步才是最关键的,起一个监听服务
	//ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("Server ready to read...")
	for {
		//AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn：接收连接
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("accept error: ", err)
			continue
		}
		//返回远端网络地址
		fmt.Println("a client connected: " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected: ", ipStr)
		conn.Close()
	}()
	//NewReader创建一个具有默认大小缓冲、从r读取的*Reader。
	//读取客户端发来的数据
	reader := bufio.NewReader(conn)
	i := 0
	for {
		//ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串
		message, err := reader.ReadString('\n') //将数据按照换行符读取
		//当有错误或者是错误是缓冲区为空的时候退出
		if err != nil || err == io.EOF {
			break
		}
		//输出客户端发来的数据
		fmt.Println(string(message))
		time.Sleep(time.Second * 3)
		msg := time.Now().String() + conn.RemoteAddr().String() + "server say hello\n"
		//传输信息字节流
		b := []byte(msg)
		conn.Write(b)
		i++
		//i>10就不再向网络传输信息了
		if i > 10 {
			break
		}
	}
}
