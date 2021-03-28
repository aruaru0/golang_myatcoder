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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	L := make([]int, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		L[i] = inf
		R[i] = -inf
	}
	for i := 0; i < N; i++ {
		x, c := getI(), getI()-1
		L[c] = min(L[c], x)
		R[c] = max(R[c], x)
	}

	l := make([]int, 0)
	r := make([]int, 0)
	for i := 0; i < N; i++ {
		if L[i] == inf {
			continue
		}
		l = append(l, L[i])
		r = append(r, R[i])
	}

	N = len(l)
	dp := make([][2]int, N+1)
	lpos, rpos := 0, 0
	for i := 0; i < N; i++ {
		if l[i] == inf {
			continue
		}
		// x ->  r -> l
		distl := 0
		if r[i] <= lpos {
			distl = lpos - l[i]
		} else {
			distl = (r[i] - lpos) + (r[i] - l[i])
		}
		distr := 0
		if r[i] <= rpos {
			distr = rpos - l[i]
		} else {
			distr = (r[i] - rpos) + (r[i] - l[i])
		}
		// out("x->r->l", distl, distr)
		dp[i+1][0] = min(dp[i][0]+distl, dp[i][1]+distr)

		// x ->  l -> r
		distl = 0
		if l[i] >= lpos {
			distl = r[i] - lpos
		} else {
			distl = lpos - l[i] + r[i] - l[i]
		}
		distr = 0
		if l[i] >= rpos {
			distr = r[i] - rpos
		} else {
			distr = rpos - l[i] + r[i] - l[i]
		}
		// out(lpos, rpos, "x->l->r", distl, distr, "LR", l[i], r[i])
		dp[i+1][1] = min(dp[i][0]+distl, dp[i][1]+distr)

		lpos = l[i]
		rpos = r[i]
		// out(dp[pos], lpos, rpos)
	}

	// out(dp)
	ans := min(dp[N][0]+abs(lpos), dp[N][1]+abs(rpos))
	out(ans)
}
