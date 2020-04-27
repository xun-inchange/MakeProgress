package main

import "fmt"

//选择排序

func main() {
	var i []int = []int{8, 3, 2, 9, 4, 6, 10, 0}
	//selectSort(i)
	//fmt.Println("选择排序正序结果：", i)
	RSelectSort(i)
	fmt.Println("选择排序倒序结果: ", i)
}
func selectSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func RSelectSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i; j < len(a); j++ {
			if a[j] > a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}
