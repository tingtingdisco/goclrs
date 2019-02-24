package main

import (
	"fmt"

	"github.com/tingtingdisco/goclrs/sort"
)

func main() {
	arr := []int{4, 2, 5, 1, 3, -1}
	sort.MergeSort(arr, 0, len(arr)-1)
	sortedArr := []int{-1, 1, 2, 3, 4, 5}

	fmt.Println(arr, sortedArr)
}
