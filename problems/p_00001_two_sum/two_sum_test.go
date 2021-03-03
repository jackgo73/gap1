package p_00001_two_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_twoSumS1(t *testing.T) {
	res := twoSumS1([]int{2, 7, 11, 15}, 9)
	require.Equal(t, res, []int{0,1})
}

func Test_twoSumS2(t *testing.T) {
	res := twoSumS2([]int{2, 7, 11, 15}, 9)
	require.Equal(t, res, []int{0,1})
}

func Benchmark_twoSumS1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		twoSumS1([]int{2, 7, 11, 15}, 9)
	}
}

func Benchmark_twoSumS2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		twoSumS2([]int{2, 7, 11, 15}, 9)
	}
}