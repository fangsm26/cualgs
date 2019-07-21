package algs

import (
	"testing"
)

type findMedian_s struct {
	nums1  []int
	nums2  []int
	median float64
}

var findMedianTests = []findMedian_s{
	{[]int{1}, []int{2}, 1.5},
	{[]int{1, 2, 3}, []int{4, 5, 6}, 3.5},
	{[]int{1, 2, 3}, []int{4, 5}, 3.0},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for i := 0; i < len(findMedianTests); i++ {
		e := findMedianTests[i]
		if f := FindMedianSortedArrays(e.nums1, e.nums2); e.median != f {
			t.Errorf("FindMedianSortedArrays(%v, %v) = %g, want %g", e.nums1, e.nums2, f, e.median)
		}
	}
}
