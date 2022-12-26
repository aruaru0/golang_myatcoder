package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

func pow(p, n int) int {
	ret := 1
	x := p
	for n != 0 {
		if n%2 == 1 {
			ret *= x
		}
		n /= 2
		x = x * x
	}
	return ret
}

type edge struct {
	to, idx int
}

var N int
var node [][]edge
var depth []int
var cnt []int
var bit int

// sからeに向かうために使った辺を調べる
func dfs(s, e, p int) bool {
	if s == e {
		return true
	}
	ret := false
	for _, v := range node[s] {
		if v.to == p {
			continue
		}
		if dfs(v.to, e, s) {
			// out(s, "->", v)
			bit |= 1 << v.idx
			ret = true
		}
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	node = make([][]edge, N)
	for i := 0; i < N-1; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], edge{b, i})
		node[b] = append(node[b], edge{a, i})
	}

	M := getI()
	u := make([]int, M)
	v := make([]int, M)
	b := make([]int, M)
	for i := 0; i < M; i++ {
		u[i], v[i] = getI()-1, getI()-1
		bit = 0
		dfs(u[i], v[i], -1)
		// uからvに行くために使った辺をビット列として格納する
		b[i] = bit
	}

	// すべての組み合わせの数
	ans := pow(2, N-1)
	m := 1 << M
	for i := 1; i < m; i++ { // 全パターン網羅
		bit := 0
		for j := 0; j < M; j++ {
			if (i>>j)%2 == 1 {
				bit |= b[j]
			}
		}
		// bitの立っている数を調べる
		c := bits.OnesCount(uint(bit))
		if bits.OnesCount(uint(i))%2 == 0 { // 包除原理
			ans += pow(2, N-1-c)
		} else {
			ans -= pow(2, N-1-c)
		}

	}
	out(ans)
}
