package p_00073_set_zeroes

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}


func setZeroes1(matrix [][]int)  {
	rows := len(matrix)
	cols := len(matrix[0])

	rmap := map[int]bool{}
	lmap := map[int]bool{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				rmap[r] = true
				lmap[c] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if rmap[i] || lmap[j] {
				r[j] = 0
			}
		}
	}
}