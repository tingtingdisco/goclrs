package array

import (
	"math"
)

// FindMaxCrossingSubarray FindMaxCrossingSubarray
func FindMaxCrossingSubarray(arr []int, lo, mid, hi int) (int, int, int) {

	leftSum := math.MinInt64
	sum := 0
	var maxLeft int

	for i := mid; i >= lo; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := math.MinInt64
	sum = 0
	var maxRight int

	for i := mid + 1; i <= hi; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}

	return maxLeft, maxRight, leftSum + rightSum

}

// FindMaxSubarray FindMaxSubarray
func FindMaxSubarray(arr []int, lo, hi int) (int, int, int) {
	if lo == hi {
		return lo, lo, arr[lo]
	}
	mid := (lo + hi) / 2

	crossLow, crossHigh, crossSum := FindMaxCrossingSubarray(arr, lo, mid, hi)
	leftLow, leftHigh, leftSum := FindMaxSubarray(arr, lo, mid)
	rightLow, rightHigh, rightSum := FindMaxSubarray(arr, mid+1, hi)

	if leftSum >= crossSum && leftSum >= rightSum {
		return leftLow, leftHigh, leftSum
	}

	if rightSum >= crossSum && rightSum >= leftSum {
		return rightLow, rightHigh, rightSum
	}

	return crossLow, crossHigh, crossSum

}
