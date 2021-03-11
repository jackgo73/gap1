package p_00056_merge

import "sort"

// 排序后总结成三种关系
//     --------       --------        --------
//     ----------       ----------              ----------

//     --------       --------
//     ---                  ---

func merge(intervals [][]int) [][]int {
	res := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if prev[1] >= curr[0] {
			//  [1,10] [9,12] = [1,12]
			//  [1,10] [2,3]  = [1,10]
			prev[1] = max(prev[1], curr[1])
		} else {
			//  [1, 10] [11, 14]
			res = append(res, prev)
			prev = curr
		}
	}
	res = append(res, prev)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
