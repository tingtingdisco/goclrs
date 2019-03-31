package main

import (
	"fmt"

	"github.com/tingtingdisco/goclrs/sort"
)

func main() {
	arr := []int{-4, -20, 110, -5, -8, -10, -1, -3, -1}
	sort.HeapSort(arr)
	fmt.Println(arr)
}
