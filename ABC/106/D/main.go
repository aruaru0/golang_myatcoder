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

func main() {
	sc.Split(bufio.ScanWords)
	N, M, Q := getInt(), getInt(), getInt()
	n := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		n[i] = make([]int, N+1)
	}
	for i := 0; i < M; i++ {
		l, r := getInt(), getInt()
		n[l][r]++
	}
	// for i := 0; i <= N; i++ {
	// 	out(n[i])
	// }
	// out("----")
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			n[i][j] += n[i][j-1]
		}
	}
	// for i := 0; i <= N; i++ {
	// 	out(n[i])
	// }
	var dp [550][550]int
	var calced [550][550]bool

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < Q; i++ {
		p, q := getInt(), getInt()
		if calced[p][q] == true {
			fmt.Fprintln(w, dp[p][q])
			continue
		}
		ans := 0
		// out("pq", p, q)
		for j := p; j <= q; j++ {
			// out(j, j, n[j][p-1], n[j][q])
			ans += n[j][q] - n[j][p-1]
		}
		// out("-----", ans)
		fmt.Fprintln(w, ans)
		calced[p][q] = true
		dp[p][q] = ans
	}
}
