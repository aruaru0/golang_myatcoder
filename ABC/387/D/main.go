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

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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
	d    int // direction 0, 1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)

	sx, sy := 0, 0
	gx, gy := 0, 0
	for i := 0; i < H; i++ {
		s[i] = getS()
		for j := 0; j < W; j++ {
			if 'S' == s[i][j] {
				sx, sy = i, j
			}
			if 'G' == s[i][j] {
				gx, gy = i, j
			}
		}
	}

	const inf = int(1e18)
	dist := make([][][2]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([][2]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = [2]int{inf, inf}
		}
	}

	dist[sx][sy][0] = 0
	dist[sx][sy][1] = 0
	q := make([]pos, 0)
	q = append(q, pos{sx, sy, 0})
	q = append(q, pos{sx, sy, 1})

	dx := [][2]int{{-1, 1}, {0, 0}}
	dy := [][2]int{{0, 0}, {-1, 1}}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		nd := 0
		if cur.d == 0 {
			nd = 1
		}
		for i := 0; i < 2; i++ {
			px, py := cur.x+dx[nd][i], cur.y+dy[nd][i]
			if px < 0 || px >= H || py < 0 || py >= W {
				continue
			}
			if s[px][py] == '#' {
				continue
			}
			if dist[px][py][nd] == inf {
				dist[px][py][nd] = dist[cur.x][cur.y][cur.d] + 1
				q = append(q, pos{px, py, nd})
			}
		}
	}

	// out(gx, gy)
	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }

	d := min(dist[gx][gy][0], dist[gx][gy][1])

	if d == inf {
		out(-1)
	} else {
		out(d)
	}
}
