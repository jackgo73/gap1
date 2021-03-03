package p_01232_check_straight_line

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkStraightLine__On__mj(t *testing.T) {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	res := checkStraightLine__On__mj(coordinates)
	assert.True(t, res)
}
