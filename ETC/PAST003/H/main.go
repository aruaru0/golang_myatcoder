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

const inf = 1001001001001001

func walk(x int, t []int, h []int) ([]int, int) {
	ret := []int{0}
	if h[x] == 1 {
		ret[0] += t[2]
	}
	ret[0] += t[0]
	return ret, x + 1
}

func sjump(x int, t []int, h []int) ([]int, int) {
	ret := []int{0, 0}
	if h[x] == 1 {
		ret[0] += t[2]
	}
	ret[0] += t[0]/2 + t[1]/2
	ret[1] += ret[0] + t[0]/2 + t[1]/2

	return ret, x + 2
}

func ljump(x int, t []int, h []int) ([]int, int) {
	ret := []int{0, 0, 0, 0}
	if h[x] == 1 {
		ret[0] += t[2]
	}
	ret[0] += t[0]/2 + t[1]/2
	ret[1] += ret[0] + t[1]
	ret[2] += ret[1] + t[1]
	ret[3] += ret[2] + t[0]/2 + t[1]/2
	return ret, x + 4
}

func main() {
	sc.Split(bufio.ScanWords)
	N, L := getInt(), getInt()
	h := make([]int, 110000)
	for i := 0; i < N; i++ {
		x := getInt()
		h[x] = 1
	}
	T1, T2, T3 := getInt(), getInt(), getInt()

	var dp [110000]int
	for i := 0; i < 110000; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 0; i < L; i++ {
		if h[i] == 1 {
			dp[i+1] = min(dp[i+1], dp[i]+T1+T3)
			dp[i+2] = min(dp[i+2], dp[i]+T1+T2+T3)
			dp[i+4] = min(dp[i+4], dp[i]+T1+3*T2+T3)
		} else {
			dp[i+1] = min(dp[i+1], dp[i]+T1)
			dp[i+2] = min(dp[i+2], dp[i]+T1+T2)
			dp[i+4] = min(dp[i+4], dp[i]+T1+3*T2)
		}
	}
	// out(dp[:L+1])
	ans := dp[L]
	// ジャンプしてＬを追加
	if h[L-1] == 1 {
		ans = min(ans, dp[L-1]+T1/2+T2/2+T3)
	} else {
		ans = min(ans, dp[L-1]+T1/2+T2/2)
	}
	if L-2 >= 0 && h[L-2] == 1 {
		ans = min(ans, dp[L-2]+T1/2+T2+T2/2+T3)
	} else {
		ans = min(ans, dp[L-2]+T1/2+T2+T2/2)
	}
	if L-3 >= 0 {
		if h[L-3] == 1 {
			ans = min(ans, dp[L-3]+T1/2+2*T2+T2/2+T3)
		} else {
			ans = min(ans, dp[L-3]+T1/2+2*T2+T2/2)
		}
	}
	out(ans)
}
