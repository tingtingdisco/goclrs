package sort

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
		q := (p + r + 1) / 2
		MergeSort(arr, p, q-1)
		MergeSort(arr, q, r)
		merge(arr, p, q, r)
	}
}

func mergeLR(left, right []int) []int {
	merged := make([]int, len(left)+len(right))

	if len(left) == 0 {
		return right
	}

	if len(right) == 0 {
		return left
	}

	i, j := 0, 0

	for k := range merged {
		if i >= len(left) && j < len(right) {
			merged[k] = right[j]
			j++
		} else if j >= len(right) && i < len(left) {
			merged[k] = left[i]
			i++
		} else if left[i] > right[j] {
			merged[k] = right[j]
			j++
		} else {
			merged[k] = left[i]
			i++
		}
	}

	return merged
}

// MergeSortArr sort whole arr
func MergeSortArr(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	return mergeLR(MergeSortArr(arr[:mid]), MergeSortArr(arr[mid:]))
}
