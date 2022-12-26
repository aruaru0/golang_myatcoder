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

func max32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

var dp [][]int32
var maxL int

func makeTbl(N int, v, w []int) {
	N = min(N, len(v))
	dp = make([][]int32, N+1)
	for n := 0; n <= N; n++ {
		dp[n] = make([]int32, maxL+1)
	}
	// out(v, w, N)

	for n := 1; n < N; n++ {
		prev := n / 2
		// out("cur", n, "prev", prev)
		for i := 0; i <= maxL; i++ {
			if w[n] <= i {
				dp[n][i] = max32(dp[prev][i], dp[prev][i-w[n]]+int32(v[n]))
			} else {
				dp[n][i] = dp[prev][i]
			}
		}
		// out(n, dp[n][:10])
	}
}

func calc(s, L int, v, w []int) int {
	N := len(v)
	n := 1 << uint(N)
	ret := 0
L0:
	for i := 0; i < n; i++ {
		sumw := 0
		sumv := 0
		for j := 0; j < N; j++ {
			if (i>>j)%2 == 1 {
				sumw += w[j]
				sumv += v[j]
			}
			if sumw > L {
				continue L0
			}
		}
		// out(strconv.FormatInt(int64(i), 2), sumw, sumv, L)
		rest := L - sumw
		sumv += int(dp[s][rest])
		ret = max(ret, sumv)
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	v := make([]int, N+1)
	w := make([]int, N+1)
	for i := 1; i <= N; i++ {
		v[i], w[i] = getInt(), getInt()
	}

	Q := getInt()
	AV := make([]int, Q)
	AL := make([]int, Q)

	for i := 0; i < Q; i++ {
		AV[i], AL[i] = getInt(), getInt()
		maxL = max(maxL, AL[i])
	}

	const lv = 2048
	makeTbl(lv, v, w)

	wb := bufio.NewWriter(os.Stdout)
	defer wb.Flush()
	for i := 0; i < Q; i++ {
		L := AL[i]
		V := AV[i]
		// out(V, L, dp[V][:L])
		if V < lv {
			fmt.Fprintln(wb, dp[V][L])
			continue
		}
		val := make([]int, 0, 16)
		weight := make([]int, 0, 16)
		for V >= lv {
			val = append(val, v[V])
			weight = append(weight, w[V])
			V /= 2
		}
		// out(val, weight, V, L)
		ret := calc(V, L, val, weight)
		fmt.Fprintln(wb, ret)
	}
}
