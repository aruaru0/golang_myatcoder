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

type ITEM struct {
	p, u, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, X, K := getI(), getI(), getI()
	items := make([]ITEM, N)
	for i := 0; i < N; i++ {
		items[i] = ITEM{getI(), getI(), getI()}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].c == items[j].c {
			if items[i].p == items[j].p {
				return items[i].u < items[j].u
			}
			return items[i].p < items[j].p
		}
		return items[i].c < items[j].c
	})

	dp := make([][2]int, X+1)

	pre_c := -1
	for _, e := range items {
		if pre_c != e.c {
			for i := 0; i < X+1; i++ {
				chmax(&dp[i][0], dp[i][1])
			}
			pre_c = e.c
		}
		old := make([][2]int, X+1)
		dp, old = old, dp
		for i := 0; i < X+1; i++ {
			for j := 0; j < 2; j++ {
				chmax(&dp[i][j], old[i][j])
			}
			ni := i + e.p
			if ni <= X {
				chmax(&dp[ni][1], old[i][1]+e.u)
				chmax(&dp[ni][1], old[i][0]+e.u+K)
			}
		}
	}

	out(dp[X][1])

}
