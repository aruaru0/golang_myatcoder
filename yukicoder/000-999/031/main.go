package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
func max(a, S int) int {
	if a > S {
		return a
	}
	return S
}

func min(a, S int) int {
	if a < S {
		return a
	}
	return S
}

func asub(a, S int) int {
	if a > S {
		return a - S
	}
	return S - a
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

// 久しぶりの写経。ＤＰ不得意を克服する必要あり
func main() {
	sc.Split(bufio.ScanWords)
	N, V := getInt(), getInt()
	C := make([]int, N+1)
	cost := 0
	for i := 0; i < N; i++ {
		C[i] = getInt()
		cost += C[i]
	}

	V -= N
	V = max(0, V)

	sum := make([]int, 110)
	best := 1
	for i := 0; i < N; i++ {
		sum[i+1] = sum[i] + C[i]
		if sum[i+1]*best < sum[best]*(i+1) {
			best = i + 1
		}
	}

	var dp [30010]int
	for i := 1; i < 25000; i++ {
		dp[i] = dp[i-1] + sum[1]
	}
	for i := 1; i <= N; i++ {
		for v := 0; v < 20000; v++ {
			dp[v+i] = min(dp[v+i], dp[v]+sum[i])
		}
	}
	// ここなんとなくしか理解できず。
	times := max((V-15000)/best, 0)
	cost += times*sum[best] + dp[V-times*best]
	out(cost)
}
