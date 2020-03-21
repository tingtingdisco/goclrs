package sort

import (
	"fmt"
	"math"
)

func Merge(arr []int, p, q, r int) {
	L := arr[p:q]
	R := arr[q : r+1]

	L = append(L, 1000000)
	R = append(R, 1000000)

	fmt.Println(L, R)

	i, j := 0, 0

	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
	}

}

func MergeSort(arr []int, p, r int) {
	if p < r {
		q := int(math.Floor((float64(p+r) / 2)))
		MergeSort(arr, p, q)
		MergeSort(arr, q+1, r)
		Merge(arr, p, q, r)
	}
}
