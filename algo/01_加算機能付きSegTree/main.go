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

//　セグメント木
const inf = 1100100100

type segtree struct {
	n int
	v []int
	a []int // 加算値
}

func (s *segtree) init(n int) {
	N := 1
	for n > N {
		N *= 2
	}
	s.v = make([]int, N*2)
	s.a = make([]int, N*2)
	s.n = N
}

func (s *segtree) update() {
	K := s.n
	for K > 0 {
		K /= 2
		for i := K; i < 2*K; i++ {
			s.v[i] = s.compare(s.v[i*2], s.v[i*2+1])
		}
	}
}

func (s *segtree) lazyCalc(k int) {
	s.v[k] += s.a[k]
	if k < s.n {
		s.a[2*k] += s.a[k]
		s.a[2*k+1] += s.a[k]
	}
	s.a[k] = 0
}

func (s *segtree) compare(p, q int) int {
	if p > q {
		return q
	}
	return p
}

func (s *segtree) setSub(a, b, k, l, r, x int) {
	if b <= l || r <= a {
		return
	}
	if a <= l && r <= b {
		s.a[k] += x
		s.lazyCalc(k)
	} else {
		s.lazyCalc(k)
		s.setSub(a, b, k*2, l, (l+r)/2, x)
		s.setSub(a, b, k*2+1, (l+r)/2, r, x)
		s.v[k] = s.compare(s.v[2*k], s.v[2*k+1])
	}
}

// set v[k] = val
func (s *segtree) set(a, b, x int) {
	s.setSub(a, b, 1, 0, s.n, x)
}

// querySub
func (s *segtree) querySub(a, b, k, l, r int) int {
	s.lazyCalc(k)
	if b <= l || r <= a {
		return inf
	}
	if a <= l && r <= b {
		return s.v[k] + s.a[k]
	}
	return s.compare(
		s.querySub(a, b, k*2, l, (l+r)/2),
		s.querySub(a, b, k*2+1, (l+r)/2, r))
}

// query([a,b) )
func (s *segtree) query(a, b int) int {
	return s.querySub(a, b, 1, 0, s.n)
}

// 連続領域加算機能つきのセグメント木
//   set(a,b,x)でエリア全体に値が設定できる
//   つけたしができないので注意、a=i, b=i+1の値にマイナス設定
func main() {
	sc.Split(bufio.ScanWords)

	var s segtree

	N, Q := getInt(), getInt()
	s.init(N)

	for i := 0; i < Q; i++ {
		l, r := getInt()-1, getInt()
		s.set(l, r, 1)
	}
	for j := 0; j < N; j++ {
		fmt.Print(s.query(j, j+1) % 2)
	}
	out()
}
