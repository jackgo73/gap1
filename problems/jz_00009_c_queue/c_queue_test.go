package jz_00009_c_queue

import (
	"fmt"
	"testing"
)

func TestCQueue_AppendTail(t *testing.T) {
	c := Constructor()
	c.AppendTail(5)
	c.AppendTail(2)
	r1 := c.DeleteHead()
	r2 := c.DeleteHead()

	fmt.Println(r1, r2)
}