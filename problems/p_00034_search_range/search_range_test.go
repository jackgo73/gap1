package p_00034_search_range

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test__searchRange__odd(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	result := []int{3, 4}
	require.Equal(t, searchRange(nums, target), result)
}

func Test__searchRange__even(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10, 10, 17, 20}
	target := 10
	result := []int{5, 6}
	require.Equal(t, searchRange(nums, target), result)
}
