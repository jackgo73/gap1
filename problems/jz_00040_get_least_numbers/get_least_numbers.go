package jz_00040_get_least_numbers

import (
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})
	return arr[:k]
}