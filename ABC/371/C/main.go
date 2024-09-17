package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

type Edge struct {
	from, to int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	Mg := getI()
	edgeG := make(map[Edge]bool)
	for i := 0; i < Mg; i++ {
		u, v := getI()-1, getI()-1
		if u > v {
			u, v = v, u
		}
		edgeG[Edge{u, v}] = true
	}
	Mh := getI()
	edgeH := make(map[Edge]bool)
	for i := 0; i < Mh; i++ {
		u, v := getI()-1, getI()-1
		if u > v {
			u, v = v, u
		}
		edgeH[Edge{u, v}] = true
	}
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
	}
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			x := getI()
			a[i][j] = x
			a[j][i] = x
		}
	}
	// for i := 0; i < N; i++ {
	// 	out(a[i])
	// }

	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = i
	}
	ans := int(1e18)
	for {
		// out("p = ", p)
		cost := 0
		used := make(map[Edge]bool)
		for e := range edgeG {
			u, v := p[e.from], p[e.to]
			if u > v {
				u, v = v, u
			}
			if edgeH[Edge{u, v}] {
				used[Edge{u, v}] = true
			} else {
				cost += a[u][v]
			}
			// out(u, v, e, used)
		}
		for e := range edgeH {
			if used[e] == false {
				cost += a[e.from][e.to]
			}
		}
		// out(used, cost)
		ans = min(ans, cost)
		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
	}
	out(ans)
}
