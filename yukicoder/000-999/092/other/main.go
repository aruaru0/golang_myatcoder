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

type pair struct {
	t, c int
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M, K := getInt(), getInt(), getInt()
	node := make([][]pair, N)
	for i := 0; i < M; i++ {
		f, t, c := getInt()-1, getInt()-1, getInt()
		node[f] = append(node[f], pair{t, c})
		node[t] = append(node[t], pair{f, c})
	}
	d := getInts(K)
	dp := make([][]bool, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]bool, N)
	}
	for j := 0; j < N; j++ {
		dp[0][j] = true
	}
	for k := 1; k <= K; k++ {
		for n := 0; n < N; n++ {
			for _, e := range node[n] {
				if e.c == d[k-1] && dp[k-1][n] == true {
					dp[k][e.t] = true
				}
			}
		}
		// out(dp[k])
	}
	cnt := 0
	for i := 0; i < N; i++ {
		if dp[K][i] {
			cnt++
		}
	}
	out(cnt)
	for i := 0; i < N; i++ {
		if dp[K][i] {
			fmt.Print(i+1, " ")
		}
	}
	out()
}
