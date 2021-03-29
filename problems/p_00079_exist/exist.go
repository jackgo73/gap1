package p_00079_exist

var direct = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上下左右

func exist(board [][]byte, word string) bool {
	//深度优先
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	var dfs func(row, col, index int) bool

	dfs = func(row, col, index int) bool {
		tmp := board[row][col]
		if board[row][col] == word[index] {
			if index == len(word)-1 {
				return true
			}
			board[row][col] = ' '
			newRow, newColumn := 0, 0
			for _, v := range direct {
				newRow = row + v[0]
				newColumn = col + v[1]
				if newRow >= 0 && newRow < m && newColumn >= 0 && newColumn < n && board[newRow][newColumn] != ' ' {
					if dfs(newRow, newColumn, index+1) {
						return true
					}
				}
			}
		}
		board[row][col] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
