package main

import "fmt"

//归并排序
//使用递归进行归并排序
func main() {
	var i []int = []int{8, 3, 2, 9, 4, 6, 10, 0}
	i = mergeSort(i)
	fmt.Println("归并排序结果是:", i)
}

func mergeSort(a []int) []int {
	//递归结束条件:切片长度为1的时候
	if len(a) < 2 {
		return a
	}
	left := a[:len(a)/2]
	right := a[len(a)/2:]
	//继续递归下去
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var resultSlice []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			resultSlice = append(resultSlice, left[0])
			left = left[1:]
		} else {
			resultSlice = append(resultSlice, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		resultSlice = append(resultSlice, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		resultSlice = append(resultSlice, right[0])
		right = right[1:]
	}
	return resultSlice
}
