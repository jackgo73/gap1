package p_00020_is_valid

type stack struct {
	s []byte
}

func NewStack(n int) *stack {
	return &stack{s: make([]byte, n)}
}

func (s *stack) Push(v byte) {
	s.s = append(s.s, v)
}
func (s *stack) Pop() byte {
	l := len(s.s)
	if l == 0 {
		return 0
	}
	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *stack) Top() byte {
	l := len(s.s)
	return s.s[l-1]
}

func (s *stack) Empty() bool {
	return len(s.s) == 0
}

func getRight(c byte) byte {
	switch c {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	default:
		return '?'
	}
}

func isValid(s string) bool {
	stack := NewStack(0)


	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack.Push(c)
			continue
		}
		if !stack.Empty() && (getRight(c) == stack.Top()) {
			stack.Pop()
		} else {
			return false
		}
	}
	return stack.Empty()
}
