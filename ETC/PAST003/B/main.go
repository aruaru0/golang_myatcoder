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

func calc(n, N, M int, s [][]int, sum []int) {
	ans := 0

	for i := 0; i < M; i++ {
		if s[n][i] == 1 {
			ans += max(0, N-sum[i])
		}
	}
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M, Q := getInt(), getInt(), getInt()

	s := make([][]int, N)
	sum := make([]int, M)
	for i := 0; i < N; i++ {
		s[i] = make([]int, M)
	}
	for i := 0; i < Q; i++ {
		a := getInt()
		// out(a)
		if a == 1 {
			n := getInt() - 1
			calc(n, N, M, s, sum)
		} else {
			n, m := getInt()-1, getInt()-1
			s[n][m] = 1
			sum[m]++
		}
	}

}
