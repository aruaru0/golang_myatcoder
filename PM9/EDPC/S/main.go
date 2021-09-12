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

func condAB(c bool, a, b int) int {
	if c {
		return a
	}
	return b
}

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	K := getS()
	D := getI()

	nbits := len(K)

	//  [digit][smaller][cond1]
	//  cond1 ... 3である
	dp := make([][2][110]int, nbits+1)

	dp[0][0][0] = 1
	for i := 0; i < nbits; i++ {
		d := int(K[i] - '0')
		for sm := 0; sm < 2; sm++ {
			num := 10
			if sm == 0 {
				num = d
			}
			for k := 0; k < D; k++ {
				for j := 0; j < num; j++ {
					dp[i+1][1][(k+j)%D] += dp[i][sm][k]
					dp[i+1][1][(k+j)%D] %= mod
				}
			}
		}
		//　smaller = 0  d = k[i]の処理
		for j := 0; j < D; j++ {
			dp[i+1][0][(j+d)%D] += dp[i][0][j]
			dp[i+1][0][(j+d)%D] %= mod
		}
	}
	ans := (dp[nbits][0][0] + dp[nbits][1][0]) % mod
	ans = (ans - 1 + mod) % mod

	out(ans)
}
