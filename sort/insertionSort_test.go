package sort

import (
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	arr := []int{2, 4, 5, 1, 3}
	InsertionSort(arr)
	sortedArr := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(arr, sortedArr) {
		t.Error(`sort error`)
	}
}
