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

const inf = int(1e9)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := []byte("0" + getS())
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	n := len(s)
	dp := make([][2]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i][0], dp[i][1] = inf, inf
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		if dp[i][1] != inf && s[i] == '0' {
			dp[i+1][0] = min(dp[i+1][0], dp[i][1]+1)
			dp[i+1][1] = min(dp[i+1][1], dp[i][1]+1)
		}
		if dp[i][0] != inf && s[i] == '1' {
			dp[i+1][0] = min(dp[i+1][0], dp[i][0]+1)
			dp[i+1][1] = min(dp[i+1][1], dp[i][0]+1)
		}
		if dp[i][1] != inf && s[i] == '1' {
			dp[i+1][1] = min(dp[i+1][1], dp[i][1])
		}
		if dp[i][0] != inf && s[i] == '0' {
			dp[i+1][0] = min(dp[i+1][0], dp[i][0])
		}
	}
	out(dp[n][0])
}
