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

const inf = 1001001001001001

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	s := getInts(N)
	M := getInt()
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i != j {
				dist[i][j] = inf
			}
		}
	}
	for i := 0; i < M; i++ {
		f, t, c := getInt(), getInt(), getInt()
		dist[f][t] = c
		dist[t][f] = c
	}

	// ワーシャルフロイド法
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// for i := 0; i < N; i++ {
	// 	out(dist[i])
	// }

	// out("-------------")

	ans := inf
	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			if i == j {
				continue
			}
			cost := dist[0][i] + dist[i][j] + dist[j][N-1]
			cost += s[i] + s[j]
			ans = min(ans, cost)
		}
	}
	out(ans)
}
