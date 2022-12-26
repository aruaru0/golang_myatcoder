package main

import (
	"bufio"
	"fmt"
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
	y, x, c int
}

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	var start, goal pos

	dp := make([][][2]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([][2]int, W)
		for j := 0; j < W; j++ {
			dp[i][j][0] = inf
			dp[i][j][1] = inf
			if s[i][j] == 'S' {
				start = pos{i, j, 0}
			}
			if s[i][j] == 'G' {
				goal = pos{i, j, 0}
			}
		}
	}

	q := make([]pos, 0)
	q = append(q, start)
	dp[start.y][start.x][0] = 0

	n := [2]int{8, 4}
	dx := [2][]int{{-1, 1, -2, 2, -2, 2, -1, 1}, {-1, 1, -1, 1}}
	dy := [2][]int{{-2, -2, -1, -1, 1, 1, 2, 2}, {-1, -1, 1, 1}}

	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		for i := 0; i < n[c.c]; i++ {
			px := c.x + dx[c.c][i]
			py := c.y + dy[c.c][i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			t := c.c
			if s[py][px] == 'R' {
				if t == 0 {
					t = 1
				} else {
					t = 0
				}
			}
			if dp[py][px][t] > dp[c.y][c.x][c.c]+1 {
				dp[py][px][t] = dp[c.y][c.x][c.c] + 1
				q = append(q, pos{py, px, t})
			}
		}
	}

	// for k := 0; k < 2; k++ {
	// 	fmt.Println("-----------------")
	// 	for i := 0; i < H; i++ {
	// 		for j := 0; j < W; j++ {
	// 			fmt.Print(dp[i][j][k], " ")
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	ans := min(dp[goal.y][goal.x][0], dp[goal.y][goal.x][1])
	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
