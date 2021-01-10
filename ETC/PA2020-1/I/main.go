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
	x, y, d int
}

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	sy, sx, gy, gx := getI()-1, getI()-1, getI()-1, getI()-1
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	var dp [2][2100][2100]int
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < 2; k++ {
				dp[k][i][j] = inf
			}
		}
	}
	dx := [2][2]int{{-1, 1}, {0, 0}}
	dy := [2][2]int{{0, 0}, {-1, 1}}

	q := make([]pos, 0)
	q = append(q, pos{sx, sy, 0})
	q = append(q, pos{sx, sy, 1})
	dp[0][sy][sx] = 0
	dp[1][sy][sx] = 0

	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		cost := dp[c.d][c.y][c.x]
		d := 0
		if c.d == 0 {
			d = 1
		}
		for i := 0; i < 2; i++ {
			px := c.x + dx[d][i]
			py := c.y + dy[d][i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if s[py][px] == '#' {
				continue
			}
			if dp[d][py][px] > cost+1 {
				dp[d][py][px] = cost + 1
				q = append(q, pos{px, py, d})
			}
		}
	}
	// for i := 0; i < H; i++ {
	// 	out(dp[0][i][:W])
	// }
	// out("-----")
	// for i := 0; i < H; i++ {
	// 	out(dp[1][i][:W])
	// }

	ans := min(dp[0][gy][gx], dp[1][gy][gx])

	if ans == inf {
		out(-1)
		return
	}
	out(ans)

}
