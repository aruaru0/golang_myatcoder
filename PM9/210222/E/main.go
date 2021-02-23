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
	N, M := getI(), getI()
	s := getS()

	dp := make([]int, N+2)
	for i := 0; i < N+1; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	last := 0
	for i := 0; i < N; i++ {
		if last > i {
			continue
		}
		ok := false
		for j := 1; j <= M; j++ {
			if i+j > N {
				continue
			}
			if s[i+j] != '1' {
				ok = true
				if dp[i+j] > dp[i]+1 {
					dp[i+j] = dp[i] + 1
					last = i + j
				}
			}
		}
		if !ok {
			out(-1)
			return
		}
	}
	// out(dp)

	if dp[N] == inf {
		out(-1)
		return
	}

	cnt := dp[N]
	cur := N
	ans := make([]int, 0)
	for {
		next := cur
		for j := cur - 1; j > cur-1-M; j-- {
			if j < 0 {
				break
			}
			if s[j] == '0' && dp[j] == cnt-1 {
				next = j
			}
		}
		cnt--
		ans = append(ans, cur-next)
		// out(cur, "->", next, cnt, ans)
		cur = next
		if cnt == 0 {
			break
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprint(wr, ans[i], " ")
	}
	out()
}
