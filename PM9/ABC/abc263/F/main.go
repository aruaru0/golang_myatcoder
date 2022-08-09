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

type pair struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	n := 1 << N
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = append([]int{0}, getInts(N)...)
	}

	dp := make([][]pair, n)
	for i := 0; i < n; i++ {
		dp[i] = append(dp[i], pair{0, i})
	}

	merge := func(dpl, dpr []pair) []pair {
		length := len(dpl)
		level := 0
		for 1<<level != length {
			level += 1
		}
		tmp := append(dpl, dpr...)
		left_max := dpl[0].x
		for _, e := range dpl {
			left_max = max(left_max, e.x)
		}
		right_max := dpr[0].x
		for _, e := range dpr {
			right_max = max(right_max, e.x)
		}
		// left win
		for i := 0; i < length; i++ {
			score, person := tmp[i].x, tmp[i].y
			next_score := score - C[person][level] + C[person][level+1] + right_max
			tmp[i] = pair{next_score, person}
		}

		// right win
		for i := length; i < length*2; i++ {
			score, person := tmp[i].x, tmp[i].y
			next_score := score - C[person][level] + C[person][level+1] + left_max
			tmp[i] = pair{next_score, person}
		}
		return tmp
	}

	for i := 0; i < n-1; i++ {
		dp = append(dp, merge(dp[i*2], dp[i*2+1]))
	}

	idx := len(dp) - 1
	ans := 0
	for _, e := range dp[idx] {
		ans = max(ans, e.x)
	}
	out(ans)
}
