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

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	s := getInts(N)
	t := getInts(M)

	dp := make([][]int, N+1)
	sum := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, M+1)
		sum[i] = make([]int, M+1)
	}
	res := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if s[i] == t[j] {
				if i-1 >= 0 && j-1 >= 0 {
					dp[i][j] += sum[i-1][j-1]
				}
				dp[i][j] = (dp[i][j] + 1) % mod
			}
			sum[i][j] += dp[i][j]
			if i-1 >= 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j-1 >= 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i-1 >= 0 && j-1 >= 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
			sum[i][j] %= mod
			if sum[i][j] < 0 {
				sum[i][j] += mod
			}
			res += dp[i][j]
			res %= mod
		}
	}
	out(res + 1)
}
