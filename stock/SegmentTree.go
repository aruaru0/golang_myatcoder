//
// セグメント木
//
type segTree []int

const segTreeDefault = 0

// マージ関数
func (s *segTree) segMerge(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func (s *segTree) segMerge(a,b, int) int {
	return a+b
}
*/

func segTreeInit(N int) segTree {
	size := 1
	for N > size {
		size <<= 1
	}
	s := make(segTree, size*2)
	for i := 0; i < size*2; i++ {
		s[i] = segTreeDefault
	}
	s[0] = 0
	return s
}

func (s *segTree) set(i, pos int) {
	N := len(*s)/2 + pos
	(*s)[N] = i
	for N > 1 {
		N /= 2
		(*s)[N] = s.segMerge((*s)[2*N], (*s)[2*N+1])
	}
}

func (s *segTree) setAll(a []int, n int) {
	N := len(*s) / 2
	for i := 0; i < n; i++ {
		(*s)[N+i] = a[i]
	}
	for N > 1 {
		N >>= 1
		for i := 0; i < N; i++ {
			pos := N + i
			(*s)[pos] = s.segMerge((*s)[2*pos], (*s)[2*pos+1])
		}
	}
}

func (s *segTree) querySub(l, r, index, L, R int) int {
	if R <= l || r <= L {
		return segTreeDefault
	}
	if l <= L && R <= r {
		return (*s)[index]
	}
	med := (L + R) / 2
	lval := s.querySub(l, r, index*2, L, med)
	rval := s.querySub(l, r, index*2+1, med, R)
	return s.segMerge(lval, rval)
}

// queryの区間は[l, r)です
func (s *segTree) query(l, r int) int {
	return (*s).querySub(l, r, 1, 0, len(*s)/2)
}

/*
	seg := segTreeInit(N)

	seg.setAll(a, N)
	seg.set(val, pos)
	seg.query(l, r) // [L, R)
*/
