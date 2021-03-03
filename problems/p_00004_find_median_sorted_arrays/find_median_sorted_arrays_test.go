package p_00004_find_median_sorted_arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test__findMedianSortedArrays__n__mj(t *testing.T) {
	res := findMedianSortedArrays__n__mj([]int{}, []int{1})
	require.Equal(t, res, 1.0)

	res = findMedianSortedArrays__n__mj([]int{}, []int{1, 2, 3, 4})
	require.Equal(t, res, 2.5)

	res = findMedianSortedArrays__n__mj([]int{}, []int{1, 2, 3, 4, 5})
	require.Equal(t, res, 3.0)

	res = findMedianSortedArrays__n__mj([]int{5}, []int{})
	require.Equal(t, res, 5.0)

	res = findMedianSortedArrays__n__mj([]int{}, []int{5, 7, 8, 10})
	require.Equal(t, res, 7.5)

	res = findMedianSortedArrays__n__mj([]int{}, []int{11, 15, 30, 66, 79})
	require.Equal(t, res, 30.0)

	// 1 3 4 5 7 10 11 15 18
	res = findMedianSortedArrays__n__mj([]int{1, 3, 5, 7, 10, 18}, []int{4, 11, 15})
	require.Equal(t, res, 7.0)
}
