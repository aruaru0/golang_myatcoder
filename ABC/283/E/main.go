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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}

	get := func(i int, flip bool) []int {
		res := make([]int, W)
		copy(res, a[i])
		if flip == true {
			for i := 0; i < W; i++ {
				res[i] ^= 1
			}
		}
		return res
	}

	check := func(a0 []int, a1 []int, a2 []int) bool {
		for i := 0; i < W; i++ {
			if a0[i] == a1[i] {
				continue
			}
			if a2[i] == a1[i] {
				continue
			}
			if i != 0 && (a1[i-1] == a1[i]) {
				continue
			}
			if i+1 < W && (a1[i+1] == a1[i]) {
				continue
			}
			return false
		}
		return true
	}

	dp := make([]int, 4)
	for i := 0; i < 4; i++ {
		dp[i] = inf
	}
	for s := 0; s < 4; s++ {
		a0 := make([]int, W)
		for i := 0; i < W; i++ {
			a0[i] = -1
		}
		a1 := get(0, s&2 != 0)
		a2 := get(1, s&1 != 0)
		cnt := 0
		if s&1 != 0 {
			cnt++
		}
		if s&2 != 0 {
			cnt++
		}
		if check(a0, a1, a2) {
			dp[s] = cnt
		}
	}

	for i := 2; i < H; i++ {
		p := make([]int, 4)
		for j := 0; j < 4; j++ {
			p[j] = inf
		}
		p, dp = dp, p
		for s := 0; s < 4; s++ {
			a0 := get(i-2, s&2 != 0)
			a1 := get(i-1, s&1 != 0)
			for x := 0; x < 2; x++ {
				a2 := get(i, x != 0)
				if check(a0, a1, a2) {
					ns := (s&1)<<1 | x
					dp[ns] = min(dp[ns], p[s]+x)
				}
			}
		}
	}

	ans := inf
	for s := 0; s < 4; s++ {
		a0 := get(H-2, s&2 != 0)
		a1 := get(H-1, s&1 != 0)
		a2 := make([]int, W)
		for i := 0; i < W; i++ {
			a2[i] = -1
		}
		if check(a0, a1, a2) {
			ans = min(ans, dp[s])
		}
	}

	if ans == inf {
		ans = -1
	}
	out(ans)
}
