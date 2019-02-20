package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	arr := []int{2, 4, 5, 1, 3}
	insertionSort(arr)
	sortedArr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr, sortedArr)

	if !reflect.DeepEqual(arr, sortedArr) {
		t.Error(`sort error`)
	}

}
