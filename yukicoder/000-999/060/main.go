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

const off = 600

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	x := make([]int, N)
	y := make([]int, N)
	hp := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt()+off, getInt()+off
		hp[i] = getInt()
	}
	var g [2000][2000]int
	for i := 0; i < K; i++ {
		x, y, w, h, d := getInt()+off, getInt()+off,
			getInt(), getInt(), getInt()
		g[x][y] += d
		g[x+w+1][y] -= d
		g[x][y+h+1] -= d
		g[x+w+1][y+h+1] += d
	}

	for i := 0; i < 2000; i++ {
		for j := 1; j < 2000; j++ {
			g[i][j] += g[i][j-1]
		}
	}

	for i := 0; i < 2000; i++ {
		for j := 1; j < 2000; j++ {
			g[j][i] += g[j-1][i]
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		xx := x[i]
		yy := y[i]
		if g[xx][yy] < hp[i] {
			ans += hp[i] - g[xx][yy]
		}
	}
	out(ans)
}
