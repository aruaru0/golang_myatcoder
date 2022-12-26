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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, l := getI(), getI()
	g := make([]int, l+5)
	for i := 0; i < n; i++ {
		x := getI()
		g[x] = 1
	}
	t := getInts(3)
	dp := make([]int, l+5)
	for i := 0; i < l+5; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 0; i < l; i++ {
		dp[i+1] = min(dp[i+1], t[0]+dp[i]+g[i+1]*t[2])
		dp[i+2] = min(dp[i+2], t[0]+t[1]+dp[i]+g[i+2]*t[2])
		dp[i+4] = min(dp[i+4], t[0]+3*t[1]+dp[i]+g[i+4]*t[2])
	}
	ans := dp[l]
	ans = min(ans, dp[l-1]+(t[0]+t[1])/2)
	ans = min(ans, dp[l-2]+(t[0]/2+t[1]*3/2))
	if l-3 >= 0 {
		ans = min(ans, dp[l-3]+(t[0]/2+t[1]*5/2))
	}
	out(ans)
}
