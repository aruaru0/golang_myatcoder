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

func solve() bool {
	n, m, a, b := getI(), getI(), getI(), getI()
	l := make([]int, m)
	r := make([]int, m)
	for i := 0; i < m; i++ {
		l[i], r[i] = getI()-1, getI()
	}

	if a == b {
		for i := 0; i < m; i++ {
			x := (r[i] - 1) / a * a
			if l[i] <= x {
				return false
			}
		}
		if (n-1)%a != 0 {
			return false
		}
		return true
	}

	dp := make([]int, b)
	dp[0] = 1
	i := 0
	l = append(l, n)
	r = append(r, n)
	m++
	for j := 0; j < m; j++ {
		w := l[j] - i - 1
		if w > a*a {
			allZero := true
			for _, v := range dp {
				if v != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				return false
			}
			for k := 0; k < b; k++ {
				dp[k] = 1
			}
		} else {
			for k := 0; k < w; k++ {
				dp = append([]int{0}, dp...)
				for x := a; x <= b; x++ {
					dp[0] = dp[0] | dp[x]
				}
				dp = dp[:b]
			}
		}

		w = r[j] - l[j]
		if w >= b {
			return false
		}
		for k := 0; k < w; k++ {
			dp = append([]int{0}, dp[:b-1]...)
		}

		i = r[j] - 1
	}

	return dp[0] == 1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	if solve() {
		out("Yes")
	} else {
		out("No")
	}

}
