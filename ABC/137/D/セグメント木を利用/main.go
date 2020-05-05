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

// Data :
type Data struct {
	day, pay int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	if p[i].pay == p[j].pay {
		return p[i].day > p[j].day
	}
	return p[i].pay > p[j].pay
}

//　セグメント木
const inf = -1

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
	if p < q {
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

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()

	J := make(Datas, N)
	n := make([]int, N)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		J[i] = Data{a, b}
		n[i] = i
	}

	sort.Sort(J)
	// out(J)
	var s segtree
	s.init(N)
	for i := 0; i < M; i++ {
		s.set(i, i+1, i+1)
	}
	// out(s)
	// out(J)
	idx := 0
	ans := 0
L0:
	for i := 1; i <= M; i++ {
		for {
			if idx == len(J) {
				break L0
			}
			day := J[idx].day
			ret := s.query(0, M-day+1)
			if ret > 0 {
				s.set(ret-1, ret, -ret)
				ans += J[idx].pay
				// out(ret, J[idx])
				idx++
				break
			}
			idx++
		}
		// out(s)
		// out(day, M, M-day, "ret", ret)
		// out(ret, i, M)
	}
	out(ans)
}
