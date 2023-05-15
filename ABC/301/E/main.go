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

type pos struct {
	x, y int
}

const inf = int(1e18)

// 解き方は、即思いついた（18個というのでbitDPだなと）
// 実装が面倒だったので、今回はコピペで対処
// このレベルの実装をサクッとできるようにならないといけない！！
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	h, w, t := getI(), getI(), getI()
	a := make([]string, h)
	var s, g pos
	oka := []pos{}
	for i := 0; i < h; i++ {
		a[i] = getS()
		for j := 0; j < w; j++ {
			if a[i][j] == 'S' {
				s = pos{i, j}
			} else if a[i][j] == 'G' {
				g = pos{i, j}
			} else if a[i][j] == 'o' {
				oka = append(oka, pos{i, j})
			}
		}
	}
	// okaにsとgも足しておく（全てお
	oka = append(oka, s, g)

	// dist[i][x][y] = oka[i]からマス(x,y)の距離
	size := len(oka)
	dist := make([][][]int, size)

	for i := 0; i < size; i++ {
		ox, oy := oka[i].x, oka[i].y
		dist[i] = make([][]int, h)
		for x := 0; x < h; x++ {
			dist[i][x] = make([]int, w)
			for y := 0; y < w; y++ {
				dist[i][x][y] = inf
			}
		}
		// 幅優先探索でお菓子からの距離を探索
		dist[i][ox][oy] = 0
		list := []pos{{ox, oy}}
		dx := []int{-1, 1, 0, 0}
		dy := []int{0, 0, -1, 1}
		for len(list) > 0 {
			mas := list[0]
			list = list[1:]
			for k := 0; k < 4; k++ {
				if mas.x+dx[k] < 0 || mas.x+dx[k] >= h ||
					mas.y+dy[k] < 0 || mas.y+dy[k] >= w ||
					a[mas.x+dx[k]][mas.y+dy[k]] == '#' {
					continue
				}
				if dist[i][mas.x+dx[k]][mas.y+dy[k]] > dist[i][mas.x][mas.y]+1 {
					dist[i][mas.x+dx[k]][mas.y+dy[k]] = dist[i][mas.x][mas.y] + 1
					if dist[i][mas.x+dx[k]][mas.y+dy[k]] < t {
						list = append(list, pos{mas.x + dx[k], mas.y + dy[k]})
					}
				}
			}
		}
	}

	// bitDPでいくつお菓子ゲットするか解く
	// dp[bit][u] = bitで示されるお菓子をもってoka[u]に行くまでの最短移動距離
	dp := make([][]int, 1<<(size-2))
	for i := 0; i < (1 << (size - 2)); i++ {
		dp[i] = make([]int, size)
		for u := 0; u < size; u++ {
			dp[i][u] = inf
		}
	}
	// dp[0]の初期化
	for i := 0; i < size; i++ {
		dp[0][i] = dist[i][s.x][s.y]
		if i < size-2 {
			dp[1<<i][i] = dp[0][i]
		}
	}

	for i := 1; i < (1 << (size - 2)); i++ {
		for u := 0; u < size; u++ {
			for v := 0; v < size; v++ {
				if dp[i][v] <= t && dist[v][oka[u].x][oka[u].y] <= t {
					ndist := dp[i][v] + dist[v][oka[u].x][oka[u].y]
					dp[i][u] = min(dp[i][u], ndist)
					if u < size-2 {
						dp[i|(1<<u)][u] = min(dp[i|(1<<u)][u], ndist)
					}
				}
			}
		}
	}

	// Gにたどりついた中で一番おかしをゲットしたものを数える
	res := -1
	for i := 0; i < (1 << (size - 2)); i++ {
		if dp[i][size-1] <= t {
			x := 0
			for j := 0; j < size-2; j++ {
				if (i>>j)&1 == 1 {
					x++
				}
			}
			res = max(res, x)
		}
	}
	fmt.Println(res)
}
