package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

//客户端
func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:999")
	//与客户端建立连接
	//连接本地地址或远端地址，这里是显然连接远端地址，所以本地地址为空
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("client connect server error! " + err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr().String() + " client connected!")
	onMessageReceived(conn)
}

func onMessageReceived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " say hello to server...\n")
	conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		//输出客户端的信息
		fmt.Println("Reading")
		fmt.Println(string(msg))

		time.Sleep(time.Second * 2)
		fmt.Println("writing...")
		b := []byte(conn.LocalAddr().String() + "write data to server!\n")
		_, err = conn.Write(b)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
