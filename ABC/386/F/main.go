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

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	k := getI()
	s, t := getS(), getS()
	sl, tl := len(s), len(t)

	if abs(sl-tl) > k {
		out("No")
		return
	}

	const inf = int(1e9)
	dp := make([][]int, sl+1)
	for i := 0; i < sl+1; i++ {
		dp[i] = make([]int, 2*k+1)
		for j := 0; j < 2*k+1; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][k] = 0
	for i := 0; i <= sl; i++ {
		for dj := 0; dj <= 2*k; dj++ {
			j := i + dj - k
			if j < 0 {
				continue
			}
			if j > tl {
				break
			}
			if i > 0 && dj < 2*k {
				dp[i][dj] = min(dp[i][dj], dp[i-1][dj+1]+1)
			}
			if j > 0 && dj > 0 {
				dp[i][dj] = min(dp[i][dj], dp[i][dj-1]+1)
			}
			if i > 0 && j > 0 {
				add := 1
				if s[i-1] == t[j-1] {
					add = 0
				}
				dp[i][dj] = min(dp[i][dj], dp[i-1][dj]+add)
			}
		}
	}

	if dp[sl][k+tl-sl] <= k {
		out("Yes")
	} else {
		out("No")
	}
}
