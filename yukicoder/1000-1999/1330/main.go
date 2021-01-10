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

const maxi = 40
const inf = math.MaxInt64 - 100

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, P := getI(), getI(), getI()
	A := getInts(N)

	Max := make([]int, maxi)
	mem := 0
	for i := 0; i < N; i++ {
		cnt := 0
		mem = max(mem, A[i])
		for A[i]%P == 0 {
			A[i] /= P
			cnt += 1
		}
		Max[cnt] = max(Max[cnt], A[i])
	}

	dp := make([]int, maxi*maxi)
	for i := 0; i < maxi*maxi; i++ {
		dp[i] = -inf
	}
	dp[0] = 1
	for i := 0; i < maxi*maxi; i++ {
		if mem*dp[i] > M {
			out(i + 1)
			return
		}
		for j := 0; j < maxi; j++ {
			if j+1+i < maxi*maxi && dp[i+j+1] < dp[i]*Max[j] {
				dp[i+j+1] = dp[i] * Max[j]
			}
		}
	}
	out(-1)
}
