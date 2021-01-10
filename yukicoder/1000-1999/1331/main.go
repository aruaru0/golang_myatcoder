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
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	A := getInts(N)

	dp := make([]int, N)
	dp[0] = 1
	M := int(math.Sqrt(float64(N)))
	Sum := make([][]int, M+1)
	for i := 0; i < M+1; i++ {
		Sum[i] = make([]int, M+1)
	}
	for i := 0; i < N; i++ {
		if i > 0 && A[i-1] != 1 {
			dp[i] += dp[i-1]
			dp[i] %= mod
		}
		for j := 0; j < M+1; j++ {
			if j > 0 {
				dp[i] += Sum[j][i%j]
				dp[i] %= mod
			}
		}
		if A[i] <= M {
			Sum[A[i]][i%A[i]] += dp[i]
			Sum[A[i]][i%A[i]] %= mod
		} else {
			for j := 0; i < N; j++ {
				X := j*A[i] + A[i] + i
				if X >= N {
					break
				}
				dp[X] += dp[i]
				dp[X] %= mod
			}
		}
	}
	out(dp[N-1])
}
