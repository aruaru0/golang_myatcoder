package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

type dat struct {
	x, y int
}

type seg []dat

const inf = 0

func segInit(N int) seg {
	n := 1
	for n < N {
		n *= 2
	}
	n *= 2
	seg := make([]dat, n)
	for i := 1; i < n; i++ {
		seg[i] = dat{inf, inf}
	}
	return seg
}

func (s seg) getValue(p int) dat {
	n := p + len(s)/2
	return s[n]
}

func (s seg) setValue(p int, v dat) {
	n := p + len(s)/2
	s[n] = v
	n /= 2
	for n != 0 {
		if s[n*2].x > s[n*2+1].x {
			s[n] = s[n*2]
		} else {
			s[n] = s[n*2+1]
		}
		n /= 2
	}
}

func (s seg) getMinSub(a, b, l, r, idx int) dat {
	// out(a, b, l, r, idx)
	if a <= l && r <= b {
		return s[idx]
	}
	if r <= a || b <= l {
		return dat{inf, inf}
	}
	m := (l + r) / 2
	x0 := s.getMinSub(a, b, l, m, idx*2)
	x1 := s.getMinSub(a, b, m, r, idx*2+1)
	if x0.x > x1.x {
		return x0
	}
	return x1
}

func (s seg) getMin(a, b int) dat {
	ret := s.getMinSub(a, b, 0, len(s)/2, 1)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)

	// s := segInit(4)
	// s.setValue(0, dat{10, 1})
	// s.setValue(1, dat{2, 2})
	// s.setValue(2, dat{4, 3})
	// s.setValue(3, dat{1, 4})

	// out(s)

	// out(s.getMin(0, 4))
	// s.setValue(3, dat{20, 4})
	// out(s.getMin(0, 4))
	N := getInt()
	K := make([]int, N)
	T := make([][]int, N)
	for i := 0; i < N; i++ {
		K[i] = getInt()
		for j := 0; j < K[i]; j++ {
			t := getInt()
			T[i] = append(T[i], t)
		}
	}
	M := getInt()
	p := make([]int, M)
	for i := 0; i < M; i++ {
		p[i] = getInt()
	}
	// t0, t1を初期化
	t0 := segInit(N)
	t1 := segInit(N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		t0.setValue(i, dat{T[i][c[i]], i})
		c[i]++
		if c[i] < K[i] {
			t1.setValue(i, dat{T[i][c[i]], i})
			c[i]++
		} else {
			t1.setValue(i, dat{inf, i})
		}
	}
	// 取っていく処理
	for i := 0; i < M; i++ {
		// out("p", p[i])
		// out("-------")
		// out(t0[16:26])
		// out(t1[16:26])
		if p[i] == 1 {
			x := t0.getMin(0, N)
			out(x.x)
			idx := x.y
			d := t1.getValue(idx)
			t0.setValue(idx, d)
			if c[idx] < K[idx] {
				t1.setValue(idx, dat{T[idx][c[idx]], idx})
				c[idx]++
			} else {
				t1.setValue(idx, dat{inf, idx})
			}
		} else {
			x := t0.getMin(0, N)
			y := t1.getMin(0, N)
			if x.x > y.x {
				out(x.x)
				idx := x.y
				d := t1.getValue(idx)
				t0.setValue(idx, d)
				if c[idx] < K[idx] {
					t1.setValue(idx, dat{T[idx][c[idx]], idx})
					c[idx]++
				} else {
					t1.setValue(idx, dat{inf, idx})
				}
			} else {
				out(y.x)
				idx := y.y
				if c[idx] < K[idx] {
					t1.setValue(idx, dat{T[idx][c[idx]], idx})
					c[idx]++
				} else {
					t1.setValue(idx, dat{inf, idx})
				}
			}
		}
	}
}
