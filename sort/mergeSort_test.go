package sort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	arr := []int{2, 4, 5, 1, 3}
	MergeSort(arr, 0, 4)
	sortedArr := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(arr, sortedArr) {
		t.Error(`sort error`)
	}
}

func Test_mergeSortArr(t *testing.T) {
	arr := MergeSortArr([]int{2, 4, 5, 1, 3})
	sortedArr := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(arr, sortedArr) {
		t.Error(`sort error`)
	}
}
