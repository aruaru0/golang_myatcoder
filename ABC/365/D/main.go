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

// Rock-Paper-Scissors = 0, 1, 2
func cost(a, b int) int {
	if a == b {
		return 0
	}
	if a == 0 && b == 1 {
		return -inf
	}
	if a == 0 && b == 2 {
		return 1
	}
	if a == 1 && b == 0 {
		return 1
	}
	if a == 1 && b == 2 {
		return -inf
	}
	if a == 2 && b == 0 {
		return -inf
	}
	if a == 2 && b == 1 {
		return 1
	}
	return 0
}

const inf = int(1e9)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	s := getS()

	m := map[byte]int{'R': 0, 'P': 1, 'S': 2}

	dp := make([][3]int, N+1)
	for i := 0; i < N+1; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = -inf
		}
	}
	dp[0][0] = 0
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 0; i < N; i++ {
		y := m[s[i]]
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if k == j {
					continue
				}
				if cost(k, y) == -inf {
					continue
				}
				dp[i+1][k] = max(dp[i+1][k], dp[i][j]+cost(k, y))
			}
		}
	}

	out(nmax(dp[N][0], dp[N][1], dp[N][2]))
}
