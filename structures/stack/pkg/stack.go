package stack

import "slices"

type Stack struct {
	items []any
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (any, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	n := len(s.items)
	last := s.items[n-1]
	s.items[n-1] = nil
	s.items = s.items[:n-1]
	return last, true
}

func (s *Stack) Peek() (any, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	n := len(s.items)
	return s.items[n-1], true

}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Items() []any {
	return slices.Clone(s.items)
}
