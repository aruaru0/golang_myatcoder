package main

import (
	"bufio"
	"fmt"
	"os"
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

//------------------------------------------------
// セグメントツリー
//------------------------------------------------
const inf = 1001001001

type rmq struct {
	seg []int
	num []int
	n   int
}

func (s *rmq) init(n int) {
	N := 1
	for N < n {
		N *= 2
	}
	s.seg = make([]int, N*2)
	s.num = make([]int, N*2)
	for i := 0; i < N*2; i++ {
		s.seg[i] = inf
	}
	s.n = N
}

func (s *rmq) update(k, x int) {
	tmp := k
	k += s.n - 1
	s.seg[k] = x
	s.num[k] = tmp
	for k > 0 {
		k = (k - 1) / 2
		if s.seg[k*2+1] <= s.seg[k*2+2] {
			s.seg[k] = s.seg[k*2+1]
			s.num[k] = s.num[k*2+1]
		} else {
			s.seg[k] = s.seg[k*2+2]
			s.num[k] = s.num[k*2+2]
		}
	}
}

// a, b ... 範囲
func (s *rmq) querySub(a, b, k, l, r int) (int, int) {
	if r <= a || b <= l {
		return inf, 0
	}
	if a <= l && r <= b {
		return s.seg[k], s.num[k]
	}

	x0, y0 := s.querySub(a, b, k*2+1, l, (l+r)/2)
	x1, y1 := s.querySub(a, b, k*2+2, (l+r)/2, r)
	if x0 <= x1 {
		return x0, y0
	} else {
		return x1, y1
	}
}

// a, b ... 範囲 a <= x < bの範囲で検索
func (s *rmq) query(a, b int) (int, int) {
	x, y := s.querySub(a, b, 0, 0, s.n)
	return x, y
}

func limit(p, k, d, n int) int {
	ret := n - p - (k-1)*d - 1
	return (ret)
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K, D := getInt(), getInt(), getInt()
	var a rmq

	a.init(N)
	for i := 0; i < N; i++ {
		x := getInt()
		a.update(i, x)
	}
	// out(a.seg)
	// out(K, D)

	f := 0
	t := limit(f, K, D, N)
	ans := make([]int, 0)
	for i := 1; i <= K; i++ {
		if t < 0 {
			out(-1)
			return
		}
		// out("query", f, t+1)
		x, y := a.query(f, t+1)
		// out(f, t, x, y)
		ans = append(ans, x)
		f = y + D
		t = f + limit(f, K-i, D, N)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, v := range ans {
		fmt.Fprint(w, v, " ")
	}
	fmt.Fprintln(w)

}
