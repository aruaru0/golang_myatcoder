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

type house struct {
	v, t int
}

var N int
var h []house

var memo [][]int

func rec(n, t int) int {
	if n == N {
		return 0
	}
	if t > 10000 {
		return 0
	}
	if memo[n][t] != 0 {
		return memo[n][t]
	}
	ret := rec(n+1, t)
	if t < h[n].t {
		ret = max(ret, rec(n+1, t+h[n].v)+h[n].v)
	}
	memo[n][t] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	h = make([]house, N)
	for i := 0; i < N; i++ {
		h[i] = house{getInt(), getInt()}
	}
	sort.Slice(h, func(i, j int) bool {
		return h[i].t+h[i].v < h[j].t+h[j].v
	})

	var dp [200200]bool
	dp[0] = true
	for i := 0; i < N; i++ {
		for j := h[i].t - 1; j >= 0; j-- {
			dp[j+h[i].v] = dp[j+h[i].v] || dp[j]
		}
		// out(h[i], dp[:15])
	}

	ans := -1
	for i := 20000; i >= 0; i-- {
		if dp[i] {
			ans = i
			break
		}
	}
	out(ans)
	// memo = make([][]int, 10000)
	// for i := 0; i < 10000; i++ {
	// 	memo[i] = make([]int, 10000)
	// }
	// ret := rec(0, 0)
	// out(ret)
}
