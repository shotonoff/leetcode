package array

type slidingWindowEntry struct {
	i int
	v int
}

type SlidingWindow struct {
	q      []*slidingWindowEntry
	h      []*slidingWindowEntry
	hLen   int
	qLimit int
}

func (s *SlidingWindow) Push(n int) {
	e := &slidingWindowEntry{0, n}
	if len(s.q) >= s.qLimit {
		e := s.q[0]
		s.q = s.q[1:]

		s.remove(e.i)
	}

	if len(s.h) > s.hLen {
		e.i = s.hLen
		s.h[s.hLen] = e
	} else {
		e.i = s.hLen
		s.h = append(s.h, e)
	}

	s.sift(s.hLen)
	s.q = append(s.q, e)
	s.hLen++
}

func (s *SlidingWindow) remove(i int) {
	s.hLen--
	s.h[i] = s.h[s.hLen]
	s.h[i].i = i
	s.down(i)
}

func (s *SlidingWindow) Max() int {
	return s.h[0].v
}

func (s *SlidingWindow) down(i int) int {
	if i >= s.hLen {
		return i
	}

	cur := i

	left := 2*i + 1
	right := 2*i + 2

	if left < s.hLen && s.h[cur].v < s.h[left].v {
		cur = left
	}

	if right < s.hLen && s.h[cur].v < s.h[right].v {
		cur = right
	}

	if cur == i {
		return i
	}

	s.h[cur], s.h[i] = s.h[i], s.h[cur]
	s.h[cur].i, s.h[i].i = cur, i

	return s.down(cur)
}

func (s *SlidingWindow) sift(i int) int {
	if i == 0 {
		return i
	}

	p := (i - 1) / 2

	if s.h[p].v < s.h[i].v {
		s.h[p], s.h[i] = s.h[i], s.h[p]
		s.h[p].i, s.h[i].i = p, i

		return s.sift(p)
	}

	return i
}
