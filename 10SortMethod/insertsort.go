package main

import "fmt"

//插入排序(顺序):将后面的数与前面的数比较,插入自己合适的位置
//3，8，2，9，4，6，10，0
func main() {
	var i []int = []int{8, 3, 2, 9, 4, 6, 10, 0}
	insertSort(i)
	RInsertSort(i)
	fmt.Println(i)
}

//正序
func insertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break //只要不比前面的数小就跳出
			}
		}
	}
}

//倒序
func RInsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}

	}
}
