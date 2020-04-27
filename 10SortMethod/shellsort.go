package main

import "fmt"

//希尔排序
//
func main() {
	var i []int = []int{8, 3, 2, 9, 4, 6, 10, 0}
	shellSort(i)
	fmt.Println("希尔排序结果是: ", i)
}

func shellSort(a []int) {
	h := 1
	for h < len(a)/3 {
		h = 3*h + 1
	}
	for h > 0 {
		for i := h; i < len(a); i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				a[j], a[j-h] = a[j-h], a[j]
			}
		}
		h /= 3
	}
}
