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

const maxn = 2000 + 10
const D = 1800
const maxD = D*2 + 10

var d = [maxn]int{}
var g = [maxn][maxD]int{}
var n, a, b int

func no() {
	fmt.Println("No")
	os.Exit(0)
}

func solve(x int, A []int) {
	for i := 1; i <= n; i++ {
		x += d[i]
	}
	if x&1 != 0 {
		no()
	}
	x /= 2
	if x < 0 {
		no()
	}
	breakItem := 0
	totWeight := 0
	for breakItem = 1; breakItem <= n; breakItem++ {
		totWeight += d[breakItem]
		if totWeight > x {
			totWeight -= d[breakItem]
			break
		}
	}
	g := [maxn][maxD]int{}
	for i := range g[breakItem-1] {
		g[breakItem-1][i] = breakItem
	}
	path := [maxn][maxD]int{}
	if D+totWeight-x < 0 {
		no()
	}
	g[breakItem-1][D+totWeight-x] = 0
	for i := breakItem; i <= n; i++ {
		for j := 1; j <= D+D; j++ {
			g[i][j] = g[i-1][j]
		}
		for j := 1; j <= D; j++ {
			if g[i-1][j] < g[i-1][j+d[i]] {
				g[i][j+d[i]] = g[i-1][j]
				path[i][j+d[i]] = i
			}
		}
		for j := D + D; j > D; j-- {
			for k := g[i][j] + 1; k <= g[i-1][j]; k++ {
				if k < g[i][j-d[k]] {
					g[i][j-d[k]] = k
					path[i][j-d[k]] = k
				}
			}
		}
	}
	if g[n][D] == breakItem {
		no()
	}
	tot := D
	for i := 1; i <= n; i++ {
		if i < breakItem {
			A[i] = 1
		} else {
			A[i] = 0
		}
	}
	for i := n; i >= breakItem; i-- {
		for path[i][tot] != 0 {
			cur := path[i][tot]
			A[cur] ^= 1
			if cur < breakItem {
				tot += d[cur]
			} else {
				tot -= d[cur]
				break
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, a, b = getI(), getI(), getI()
	for i := 1; i <= n; i++ {
		d[i] = getI()
	}

	x, y := 0, 0
	x += a
	y += a
	x -= b
	y += b
	A := make([]int, maxn)
	B := make([]int, maxn)
	solve(x, A)
	solve(y, B)

	out("Yes")
	for i := 1; i <= n; i++ {
		if A[i] != 0 && B[i] != 0 {
			fmt.Fprint(wr, "R")
		} else if A[i] != 0 {
			fmt.Fprint(wr, "D")
		} else if B[i] != 0 {
			fmt.Fprint(wr, "U")
		} else {
			fmt.Fprint(wr, "L")
		}
	}
	out()
}
