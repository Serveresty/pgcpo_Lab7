package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 4, 5, 2, 1}

	arr = sorting(arr)

	fmt.Println(arr)
}

func sorting(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}

func isSorted(arr []int) bool {
	res1 := sort.SliceIsSorted(arr, func(p, q int) bool {
		return arr[p] < arr[q]
	})

	return res1
}
