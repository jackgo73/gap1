package jz_00009_c_queue

type Stack struct {
	data []int
}

func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *Stack) Pop() int {
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}

func (s *Stack) size() int {
	return len(s.data)
}

type CQueue struct {
	inStack  *Stack
	outStack *Stack
}

func Constructor() CQueue {
	return CQueue{
		inStack:  &Stack{data: []int{}},
		outStack: &Stack{data: []int{}},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.outStack.size() > 0 {
		return this.outStack.Pop()
	} else {
		if this.inStack.size() == 0 {
			return -1
		}
		size := this.inStack.size()
		for i := 0; i < size; i++ {
			t := this.inStack.Pop()
			this.outStack.Push(t)
		}
		return this.outStack.Pop()
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
