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

type pair struct {
	x, y int
}

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	dxdy := make([]pair, 3)
	for i := 0; i < 3; i++ {
		dxdy[i] = pair{getI(), getI()}
	}

	p := make(map[pair]bool, M)
	for i := 0; i < M; i++ {
		p[pair{getI(), getI()}] = true
	}
	var dp [301][301][301]int
	dp[0][0][0] = 1
	for i := 0; i < N; i++ {
		for a := 0; a <= i; a++ {
			for b := 0; b <= i; b++ {
				x := a*dxdy[0].x + b*dxdy[1].x + (i-a-b)*dxdy[2].x
				y := a*dxdy[0].y + b*dxdy[1].y + (i-a-b)*dxdy[2].y
				for k := 0; k < 3; k++ {
					dx, dy := dxdy[k].x, dxdy[k].y
					if p[pair{x + dx, y + dy}] == false {
						if k == 0 {
							dp[i+1][a+1][b] += dp[i][a][b]
							dp[i+1][a+1][b] %= mod
						} else {
							dp[i+1][a][b+1] += dp[i][a][b]
							dp[i+1][a][b+1] %= mod
						}
					}
				}
			}
		}
	}
	ans := 0
	for a := 0; a <= N; a++ {
		for b := 0; b <= N; b++ {
			ans += dp[N][a][b]
			ans %= mod
		}
	}
	out(ans)
}
