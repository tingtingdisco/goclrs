package main

import (
	"fmt"

	"github.com/tingtingdisco/goclrs/array"
)

func main() {
	arr := []int{-4, 2, -10, 5, 8, 10, 1, 3, -1}
	l, r, s := array.FindMaxSubarray(arr, 0, len(arr)-1)
	fmt.Println(l, r, s)
}
