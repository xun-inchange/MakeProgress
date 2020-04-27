package main

import "fmt"

//冒泡排序
func main() {
	var i []int = []int{8, 3, 2, 9, 4, 6, 10, 0}
	bubbleSort(i)
	fmt.Println("冒泡排序结果: ", i)
}

func bubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		fmt.Println(a)
		fmt.Println(111)
	}
}
