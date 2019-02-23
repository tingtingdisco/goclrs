package main

import (
	"fmt"

	"github.com/tingtingdisco/goclrs/sort"
)

func main() {
	arr := []int{2, 4, 5, 1, 3}
	sort.MergeSort(arr, 0, 3)
	sortedArr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr, sortedArr)
}
