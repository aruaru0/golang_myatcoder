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

func main() {
	sc.Split(bufio.ScanWords)

	N, M, K := getInt(), getInt(), getInt()
	dp0 := make([]int, 550)

	dp0[0] = 1
	for i := 0; i < N; i++ {
		dp1 := make([]int, 550)
		a := getInts(M)
		for j := 500; j >= 0; j-- {
			for k := 0; k < M; k++ {
				if j-a[k] >= 0 {
					dp1[j] += dp0[j-a[k]]
				}
			}
		}
		dp0 = dp1
		// out(dp0[:100])
	}

	ans := -1
	for i := 0; i <= K; i++ {
		if dp0[i] != 0 {
			ans = max(ans, i)
		}
	}
	if ans == -1 {
		out(-1)
		return
	}
	out(K - ans)
}
