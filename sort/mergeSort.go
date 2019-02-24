package sort

import (
	"math"
)

func merge(arr []int, p, q, r int) {
	L := make([]int, len(arr[p:q]))
	copy(L, arr[p:q])

	R := make([]int, len(arr[q:r+1]))
	copy(R, arr[q:r+1])

	L = append(L, 1000000)
	R = append(R, 1000000)

	i, j := 0, 0

	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
	}
}

// MergeSort sort arr from index p to r
func MergeSort(arr []int, p, r int) {
	if p < r {
		q := int(math.Ceil((float64(p+r) / 2)))
		MergeSort(arr, p, q-1)
		MergeSort(arr, q, r)
		merge(arr, p, q, r)
	}
}
