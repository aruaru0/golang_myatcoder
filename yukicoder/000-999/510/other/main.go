package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
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

func mulMod(A, B [4][4]int) [4][4]int {
	H := 4
	W := 4
	K := 4
	var C [4][4]int

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j]
				C[i][j] %= mod
			}
		}
	}

	return C
}

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	q := getI()

	x := make([]int, n+1)
	y := make([]int, n+1)

	e := func() S {
		return S{[4][4]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}
	}
	merger := func(a, b S) S {
		return S{mulMod(b.a, a.a)}
	}
	D := make([]S, n+1)
	for i := 0; i <= n; i++ {
		D[i].a[0][0] = 1
		D[i].a[1][3] = 1
		D[i].a[2][3] = 1
		D[i].a[3][3] = 1
	}
	seg := newSegtree(D, e, merger)

	for k := 0; k < q; k++ {
		op := getS()
		switch op {
		case "x":
			i, e := getI(), getI()
			x[i] = e
			var v [4][4]int
			v[0][0] = 1
			v[1][3] = 1
			v[2][3] = 1
			v[3][3] = 1
			v[0][2] = x[i] % mod
			v[1][1] = y[i] % mod
			v[2][1] = 2 * y[i] % mod
			v[2][2] = y[i] * y[i] % mod
			seg.Set(i, S{v})
		case "y":
			i, e := getI(), getI()
			y[i] = e
			var v [4][4]int
			v[0][0] = 1
			v[1][3] = 1
			v[2][3] = 1
			v[3][3] = 1
			v[0][2] = x[i] % mod
			v[1][1] = y[i] % mod
			v[2][1] = 2 * y[i] % mod
			v[2][2] = y[i] * y[i] % mod
			seg.Set(i, S{v})
		case "a":
			i := getI()
			ret := seg.Prod(0, i).a
			tot := 0
			for l := 0; l < 4; l++ {
				tot += ret[0][l]
				tot %= mod
			}
			out(tot)
		}
	}

	// out(seg)
	// for i := 0; i < 5; i++ {
	// 	seg.Set(i, Data(rand.Intn(50)))
	// 	//seg.Set(i, Data(i+1))
	// }
	// seg.Update()
	// out(seg)
	// ret := seg.Query(0, 4)
	// out("Query", ret)
	// seg.UpdateAt(3, 1)
	// out(seg)
	// ret = seg.Query(1, 4)
	// out("Query", ret)
	// out("Get", seg.Get(4))

}

type S struct {
	a [4][4]int
}
type E func() S
type Merger func(a, b S) S
type Compare func(v int) bool
type Segtree struct {
	n      int
	size   int
	log    int
	d      []S
	e      E
	merger Merger
}

func newSegtree(v []S, e E, m Merger) *Segtree {
	seg := new(Segtree)
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]S, 2*seg.size)
	seg.e = e
	seg.merger = m
	for i, _ := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		seg.d[seg.size+i] = v[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.Update(i)
	}
	return seg
}
func (seg *Segtree) Update(k int) {
	seg.d[k] = seg.merger(seg.d[2*k], seg.d[2*k+1])
}
func (seg *Segtree) Set(p int, x S) {
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.Update(p >> uint(i))
	}
}
func (seg *Segtree) Get(p int) S {
	return seg.d[p+seg.size]
}
func (seg *Segtree) Prod(l, r int) S {
	sml, smr := seg.e(), seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if (l & 1) == 1 {
			sml = seg.merger(sml, seg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = seg.merger(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.merger(sml, smr)
}
func (seg *Segtree) AllProd() S {
	return seg.d[1]
}

func (seg *Segtree) ceilPow2(n int) int {
	x := 0
	for (1 << uint(x)) < n {
		x++
	}
	return x
}
