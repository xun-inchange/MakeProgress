package main

import "fmt"

//快速排序

func main() {
	var i []int = []int{8, 3, 2, 9, 4, 6, 10, 0}
	i = quickSort(i, 0, len(i)-1)
	fmt.Println("快速排序的结果是: ", i)

}

func quickSort(a []int, left, right int) []int {
	if left < right {
		//划分
		partitionIndex := partition(a, left, right)
		quickSort(a, left, partitionIndex-1)
		quickSort(a, partitionIndex+1, right)
	}
	return a

}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index = index + 1
		}
	}
	swap(arr, pivot, index-1)
	fmt.Println(arr)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
