package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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
	y, x int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([]string, H)
	stx, sty := 0, 0
	gtx, gty := 0, 0
	for i := 0; i < H; i++ {
		a[i] = getS()
		for j := 0; j < W; j++ {
			if a[i][j] == 'S' {
				stx, sty = j, i
			}
			if a[i][j] == 'T' {
				gtx, gty = j, i
			}
		}
	}
	N := getI()
	r := make([]int, N)
	c := make([]int, N)
	e := make([]int, N)

	start := -1
	for i := 0; i < N; i++ {
		r[i] = getI() - 1
		c[i] = getI() - 1
		e[i] = getI()
		if r[i] == sty && c[i] == stx {
			start = i
		}
	}

	if start == -1 {
		out("No")
		return
	}

	node := make([][]int, N+1)

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	var solve func(int)
	solve = func(n int) {
		q := []pair{}
		sx, sy := c[n], r[n]
		q = append(q, pair{sy, sx})
		dist := make([][]int, H)
		for i := 0; i < H; i++ {
			dist[i] = make([]int, W)
			for j := 0; j < W; j++ {
				dist[i][j] = -1
			}
		}
		dist[sy][sx] = 0
		for len(q) != 0 {
			cur := q[0]
			q = q[1:]
			for i := 0; i < 4; i++ {
				px, py := cur.x+dx[i], cur.y+dy[i]
				if px < 0 || px >= W || py < 0 || py >= H {
					continue
				}
				if dist[py][px] != -1 {
					continue
				}
				if a[py][px] == '#' {
					continue
				}
				dist[py][px] = dist[cur.y][cur.x] + 1
				q = append(q, pair{py, px})
			}
		}
		// out("----")
		// for i := 0; i < H; i++ {
		// 	out(dist[i])
		// }
		for i := 0; i < N; i++ {
			if n == i {
				continue
			}
			if dist[r[i]][c[i]] != -1 && dist[r[i]][c[i]] <= e[n] {
				// out(n, "-->", i, r[i], c[i], dist[r[i]][c[i]], e[i])
				node[n] = append(node[n], i)
			}
		}
		if dist[gty][gtx] != -1 && dist[gty][gtx] <= e[n] {
			node[n] = append(node[n], N)
		}
	}

	for ni := 0; ni < N; ni++ {
		solve(ni)
	}

	var dfs func(int) bool
	used := make([]bool, N+1)

	dfs = func(v int) bool {
		if v == N {
			return true
		}
		used[v] = true
		for _, e := range node[v] {
			if used[e] {
				continue
			}
			if dfs(e) {
				return true
			}
		}
		return false
	}

	ret := dfs(start)

	if ret {
		out("Yes")
	} else {
		out("No")
	}

}
