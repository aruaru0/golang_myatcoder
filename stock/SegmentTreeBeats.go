// Segment Tree Beats
// - l<=i<r について、 A_i の値を min(A_i, x) に更新
// - l<=i<r について、 A_i の値を max(A_i, x) に更新
// - l<=i<r の中の A_i の最大値を求める
// - l<=i<r の中の A_i の最小値を求める
// - l<=i<r の A_i の和を求める
// - l<=i<r について、 A_i の値に x を加える
// - l<=i<r について、 A_i の値を x に更新
type segmentTree struct {
	inf               int
	n, n0             int
	maxV, smaxV, maxC []int
	minV, sminV, minC []int
	sum               []int
	len, ladd, lval   []int
}

func (s *segmentTree) updateNodeMax(k, x int) {
	s.sum[k] += (x - s.maxV[k]) * s.maxC[k]

	if s.maxV[k] == s.minV[k] {
		s.maxV[k] = x
		s.minV[k] = x
	} else if s.maxV[k] == s.sminV[k] {
		s.maxV[k] = x
		s.sminV[k] = x
	} else {
		s.maxV[k] = x
	}

	if s.lval[k] != s.inf && x < s.lval[k] {
		s.lval[k] = x
	}
}

func (s *segmentTree) updateNodeMin(k, x int) {
	s.sum[k] += (x - s.minV[k]) * s.minC[k]

	if s.maxV[k] == s.minV[k] {
		s.maxV[k] = x
		s.minV[k] = x
	} else if s.smaxV[k] == s.minV[k] {
		s.minV[k] = x
		s.smaxV[k] = x
	} else {
		s.minV[k] = x
	}

	if s.lval[k] != s.inf && s.lval[k] < x {
		s.lval[k] = x
	}
}

func (s *segmentTree) push(k int) {

	if s.n0-1 <= k {
		return
	}

	if s.lval[k] != s.inf {
		s.updateall(2*k+1, s.lval[k])
		s.updateall(2*k+2, s.lval[k])
		s.lval[k] = s.inf
		return
	}

	if s.ladd[k] != 0 {
		s.addall(2*k+1, s.ladd[k])
		s.addall(2*k+2, s.ladd[k])
		s.ladd[k] = 0
	}

	if s.maxV[k] < s.maxV[2*k+1] {
		s.updateNodeMax(2*k+1, s.maxV[k])
	}
	if s.minV[2*k+1] < s.minV[k] {
		s.updateNodeMin(2*k+1, s.minV[k])
	}

	if s.maxV[k] < s.maxV[2*k+2] {
		s.updateNodeMax(2*k+2, s.maxV[k])
	}
	if s.minV[2*k+2] < s.minV[k] {
		s.updateNodeMin(2*k+2, s.minV[k])
	}
}

func (s *segmentTree) update(k int) {
	s.sum[k] = s.sum[2*k+1] + s.sum[2*k+2]

	if s.maxV[2*k+1] < s.maxV[2*k+2] {
		s.maxV[k] = s.maxV[2*k+2]
		s.maxC[k] = s.maxC[2*k+2]
		s.smaxV[k] = max(s.maxV[2*k+1], s.smaxV[2*k+2])
	} else if s.maxV[2*k+1] > s.maxV[2*k+2] {
		s.maxV[k] = s.maxV[2*k+1]
		s.maxC[k] = s.maxC[2*k+1]
		s.smaxV[k] = max(s.smaxV[2*k+1], s.maxV[2*k+2])
	} else {
		s.maxV[k] = s.maxV[2*k+1]
		s.maxC[k] = s.maxC[2*k+1] + s.maxC[2*k+2]
		s.smaxV[k] = max(s.smaxV[2*k+1], s.smaxV[2*k+2])
	}

	if s.minV[2*k+1] < s.minV[2*k+2] {
		s.minV[k] = s.minV[2*k+1]
		s.minC[k] = s.minC[2*k+1]
		s.sminV[k] = min(s.sminV[2*k+1], s.minV[2*k+2])
	} else if s.minV[2*k+1] > s.minV[2*k+2] {
		s.minV[k] = s.minV[2*k+2]
		s.minC[k] = s.minC[2*k+2]
		s.sminV[k] = min(s.minV[2*k+1], s.sminV[2*k+2])
	} else {
		s.minV[k] = s.minV[2*k+1]
		s.minC[k] = s.minC[2*k+1] + s.minC[2*k+2]
		s.sminV[k] = min(s.sminV[2*k+1], s.sminV[2*k+2])
	}
}

func (s *segmentTree) _updateMin(x, a, b, k, l, r int) {
	if b <= l || r <= a || s.maxV[k] <= x {
		return
	}
	if a <= l && r <= b && s.smaxV[k] < x {
		s.updateNodeMax(k, x)
		return
	}

	s.push(k)
	s._updateMin(x, a, b, 2*k+1, l, (l+r)/2)
	s._updateMin(x, a, b, 2*k+2, (l+r)/2, r)
	s.update(k)
}

func (s *segmentTree) _updateMax(x, a, b, k, l, r int) {
	if b <= l || r <= a || x <= s.minV[k] {
		return
	}
	if a <= l && r <= b && x < s.sminV[k] {
		s.updateNodeMin(k, x)
		return
	}

	s.push(k)
	s._updateMax(x, a, b, 2*k+1, l, (l+r)/2)
	s._updateMax(x, a, b, 2*k+2, (l+r)/2, r)
	s.update(k)
}

func (s *segmentTree) addall(k, x int) {
	s.maxV[k] += x
	if s.smaxV[k] != -s.inf {
		s.smaxV[k] += x
	}
	s.minV[k] += x
	if s.sminV[k] != s.inf {
		s.sminV[k] += x
	}

	s.sum[k] += s.len[k] * x
	if s.lval[k] != s.inf {
		s.lval[k] += x
	} else {
		s.ladd[k] += x
	}
}

func (s *segmentTree) updateall(k, x int) {
	s.maxV[k] = x
	s.smaxV[k] = -s.inf
	s.minV[k] = x
	s.sminV[k] = s.inf
	s.maxC[k] = s.len[k]
	s.minC[k] = s.len[k]

	s.sum[k] = x * s.len[k]
	s.lval[k] = x
	s.ladd[k] = 0
}

func (s *segmentTree) _addVal(x, a, b, k, l, r int) {
	if b <= l || r <= a {
		return
	}
	if a <= l && r <= b {
		s.addall(k, x)
		return
	}

	s.push(k)
	s._addVal(x, a, b, 2*k+1, l, (l+r)/2)
	s._addVal(x, a, b, 2*k+2, (l+r)/2, r)
	s.update(k)
}

func (s *segmentTree) _updateVal(x, a, b, k, l, r int) {
	if b <= l || r <= a {
		return
	}
	if a <= l && r <= b {
		s.updateall(k, x)
		return
	}

	s.push(k)
	s._updateVal(x, a, b, 2*k+1, l, (l+r)/2)
	s._updateVal(x, a, b, 2*k+2, (l+r)/2, r)
	s.update(k)
}

func (s *segmentTree) _queryMax(a, b, k, l, r int) int {
	if b <= l || r <= a {
		return -s.inf
	}
	if a <= l && r <= b {
		return s.maxV[k]
	}
	s.push(k)
	lv := s._queryMax(a, b, 2*k+1, l, (l+r)/2)
	rv := s._queryMax(a, b, 2*k+2, (l+r)/2, r)
	return max(lv, rv)
}

func (s *segmentTree) _queryMin(a, b, k, l, r int) int {
	if b <= l || r <= a {
		return s.inf
	}
	if a <= l && r <= b {
		return s.minV[k]
	}
	s.push(k)
	lv := s._queryMin(a, b, 2*k+1, l, (l+r)/2)
	rv := s._queryMin(a, b, 2*k+2, (l+r)/2, r)
	return min(lv, rv)
}

func (s *segmentTree) _querySum(a, b, k, l, r int) int {
	if b <= l || r <= a {
		return 0
	}
	if a <= l && r <= b {
		return s.sum[k]
	}
	s.push(k)
	lv := s._querySum(a, b, 2*k+1, l, (l+r)/2)
	rv := s._querySum(a, b, 2*k+2, (l+r)/2, r)
	return lv + rv
}

func newSegTree(n, inf int, a []int) *segmentTree {
	var s segmentTree
	s.inf = inf
	s.n = n
	s.n0 = 1
	for s.n0 < s.n {
		s.n0 <<= 1
	}
	s.maxV = make([]int, s.n0*4)
	s.smaxV = make([]int, s.n0*4)
	s.maxC = make([]int, s.n0*4)
	s.minV = make([]int, s.n0*4)
	s.sminV = make([]int, s.n0*4)
	s.minC = make([]int, s.n0*4)
	s.sum = make([]int, s.n0*4)
	s.len = make([]int, s.n0*4)
	s.ladd = make([]int, s.n0*4)
	s.lval = make([]int, s.n0*4)

	for i := 0; i < 2*s.n0; i++ {
		s.ladd[i] = 0
		s.lval[i] = inf
	}
	s.len[0] = s.n0
	for i := 0; i < s.n0-1; i++ {
		s.len[2*i+1] = (s.len[i] >> 1)
		s.len[2*i+2] = (s.len[i] >> 1)
	}

	for i := 0; i < s.n; i++ {
		val := 0
		if a != nil {
			val = a[i]
		}
		s.maxV[s.n0-1+i] = val
		s.minV[s.n0-1+i] = val
		s.sum[s.n0-1+i] = val
		s.smaxV[s.n0-1+i] = -s.inf
		s.sminV[s.n0-1+i] = s.inf
		s.maxC[s.n0-1+i] = 1
		s.minC[s.n0-1+i] = 1
	}

	for i := s.n; i < s.n0; i++ {
		s.maxV[s.n0-1+i] = -s.inf
		s.smaxV[s.n0-1+i] = -s.inf
		s.minV[s.n0-1+i] = s.inf
		s.sminV[s.n0-1+i] = s.inf
		s.maxC[s.n0-1+i] = 0
		s.minC[s.n0-1+i] = 0
	}
	for i := s.n0 - 2; i >= 0; i-- {
		s.update(i)
	}
	return &s
}

// range minimize query
func (s *segmentTree) updateMin(a, b, x int) {
	s._updateMin(x, a, b, 0, 0, s.n0)
}

// range maximize query
func (s *segmentTree) updateMax(a, b, x int) {
	s._updateMax(x, a, b, 0, 0, s.n0)
}

// range add query
func (s *segmentTree) addVal(a, b, x int) {
	s._addVal(x, a, b, 0, 0, s.n0)
}

// range update query
func (s *segmentTree) updateVal(a, b, x int) {
	s._updateVal(x, a, b, 0, 0, s.n0)
}

// range minimum query
func (s *segmentTree) queryMax(a, b int) int {
	return s._queryMax(a, b, 0, 0, s.n0)
}

// range maximum query
func (s *segmentTree) queryMin(a, b int) int {
	return s._queryMin(a, b, 0, 0, s.n0)
}

// range sum query
func (s *segmentTree) querySum(a, b int) int {
	return s._querySum(a, b, 0, 0, s.n0)
}
