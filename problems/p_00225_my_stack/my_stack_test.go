package p_00225_my_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test__MyStack(t *testing.T) {
	myStack := Constructor()

	myStack.Push(1)
	myStack.Push(2)
	r := myStack.Top()
	assert.Equal(t, r, 2)
	r = myStack.Pop()
	assert.Equal(t, r, 2)
	myStack.Empty()
}
