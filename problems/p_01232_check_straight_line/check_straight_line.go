package p_01232_check_straight_line

func checkStraightLine__On__mj(coordinates [][]int) bool {
	for i := 1; i < len(coordinates)-1; i++ {
		y1 := coordinates[i][1] - coordinates[i-1][1]
		x1 := coordinates[i][0] - coordinates[i-1][0]
		y2 := coordinates[i+1][1] - coordinates[i][1]
		x2 := coordinates[i+1][0] - coordinates[i][0]
		if y1*x2 != y2*x1 {
			return false
		}
	}
	return true
}
