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

const mod int = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, D := getI(), getI()
	p := getInts(N)
	q := getInts(N)
	// dp[p][q][r] d(p,r) < D and d(q, r)の数
	dp := make([][]int, D+1)
	for i := 0; i < D+1; i++ {
		dp[i] = make([]int, D+1)
	}
	dp[0][0] = 1

	for n := 0; n < N; n++ {
		pn, qn := p[n], q[n]
		s := abs(pn - qn)
		nxt := make([][]int, D+1)
		for i := 0; i < D+1; i++ {
			nxt[i] = make([]int, D+1)
		}
		dp2 := make([][]int, D+1)
		for i := 0; i < D+1; i++ {
			dp2[i] = make([]int, D+1)
		}
		for i := 0; i < D+1; i++ {
			for j := 0; j < D+1; j++ {
				dp2[i][j] = dp[i][j]
				if i != 0 && j != D {
					dp2[i][j] += dp2[i-1][j+1]
					dp2[i][j] %= mod
				}
			}
		}
		for i := 0; i < D+1; i++ {
			for j := 0; j < D+1; j++ {
				si, sj := i, j-s
				if sj < 0 {
					si += sj
					sj = 0
				}
				if 0 <= si && si <= D && 0 <= sj && sj <= D {
					nxt[i][j] += dp2[si][sj]
					nxt[i][j] %= mod
				}
				ti := i - (s + 1)
				tj := j + 1
				if 0 <= ti && ti <= D && 0 <= tj && tj <= D {
					nxt[i][j] -= dp2[ti][tj]
					if nxt[i][j] < 0 {
						nxt[i][j] += mod
					}
					nxt[i][j] %= mod
				}
			}
		}
		dp3 := make([][]int, D+1)
		for i := 0; i < D+1; i++ {
			dp3[i] = make([]int, D+1)
		}
		for i := 0; i < D+1; i++ {
			for j := 0; j < D+1; j++ {
				dp3[i][j] = dp[i][j]
				if i != 0 && j != 0 {
					dp3[i][j] += dp3[i-1][j-1]
					dp3[i][j] %= mod
				}
				if i+1 <= D && j+s+1 <= D {
					nxt[i+1][j+s+1] += dp3[i][j]
					nxt[i+1][j+s+1] %= mod
				}
				if i+s+1 <= D && j+1 <= D {
					nxt[i+s+1][j+1] += dp3[i][j]
					nxt[i+s+1][j+1] %= mod
				}
			}
		}
		dp = nxt
	}
	tot := 0
	for _, e := range dp {
		for _, v := range e {
			tot += v
			tot %= mod
		}
	}
	out(tot)
}
