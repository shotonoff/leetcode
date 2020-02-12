package stack

// https://leetcode.com/problems/min-stack/
// 155. Min Stack
type MinStack struct {
	h []int
	m []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	l := len(s.m)
	s.h = append(s.h, x)

	if l > 0 && x > s.m[l-1] {
		x = s.m[l-1]
	}

	s.m = append(s.m, x)
}

func (s *MinStack) Pop() {
	s.h = s.h[0 : len(s.h)-1]
	s.m = s.m[0 : len(s.m)-1]
}

func (s *MinStack) Top() int {
	return s.h[len(s.h)-1]
}

func (s *MinStack) GetMin() int {
	return s.m[len(s.m)-1]
}
