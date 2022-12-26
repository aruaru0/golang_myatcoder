package main

import (
	"bufio"
	"fmt"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	a := make([]bool, 11000)
	for i := 2; i*i <= 10000; i++ {
		for j := i * 2; j <= 10000; j += i {
			a[j] = true
		}
	}
	prime := make(map[int]bool)
	for i := 2; i <= 10000; i++ {
		if a[i] == false {
			prime[i] = true
		}
	}

	M := getI()
	N := getI()
	c := getInts(N)

	dp := make([]int, 11000)
	for i := 0; i < N; i++ {
		dp[c[i]] = 1
	}
	for i := 0; i < N; i++ {
		for j := 0; j <= M; j++ {
			if j-c[i] >= 0 && dp[j-c[i]] > 0 {
				dp[j] = max(dp[j], dp[j-c[i]]+1)
			}
		}
	}

	ans := 0
	for i := 0; i <= M; i++ {
		ans = max(ans, dp[i])
	}

	for e := range prime {
		if M-e > 0 {
			ans += dp[M-e]
		}
	}
	out(ans)
}
