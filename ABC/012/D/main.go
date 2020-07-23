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

const inf = int(1e15)

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	n := make([][]int, N)
	for i := 0; i < N; i++ {
		n[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			n[i][j] = inf
		}
	}
	for i := 0; i < M; i++ {
		a, b, t := getInt()-1, getInt()-1, getInt()
		n[a][b] = t
		n[b][a] = t
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if n[i][j] > n[i][k]+n[k][j] {
					n[i][j] = n[i][k] + n[k][j]
				}
			}
		}
	}

	ans := inf
	for i := 0; i < N; i++ {
		t := 0
		for j := 0; j < N; j++ {
			t = max(t, n[i][j])
		}
		ans = min(ans, t)
	}
	out(ans)
}
