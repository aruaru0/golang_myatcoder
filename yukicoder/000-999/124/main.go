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

var w, h int
var m [][]int
var f [][][4]int
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

const inf = int(1e12)

type pair struct {
	x, y, dir, v int
}

func bsf(sx, sy int) {
	q := make([]pair, 0)
	f[sy][sx][0] = 0
	f[sy+1][sx][3] = 1
	f[sy][sx+1][1] = 1
	q = append(q, pair{sx + 1, sy, 1, m[sy][sx]})
	q = append(q, pair{sx, sy + 1, 3, m[sy][sx]})
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		v := m[cur.y][cur.x]
		// out(cur)
		for i := 0; i < 4; i++ {
			nx := cur.x + dx[i]
			ny := cur.y + dy[i]
			if nx < 0 || nx >= w || ny < 0 || ny >= h {
				continue
			}
			if cur.v == m[ny][nx] {
				continue
			}
			// out(cur, f[ny][nx][i], f[cur.y][cur.x][cur.dir])
			if f[ny][nx][i] <= f[cur.y][cur.x][cur.dir]+1 {
				continue
			}
			// out(nx, ny)
			if (v > cur.v && v > m[ny][nx]) || (v < cur.v && v < m[ny][nx]) {
				f[ny][nx][i] = f[cur.y][cur.x][cur.dir] + 1
				q = append(q, pair{nx, ny, i, v})
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	w, h = getInt(), getInt()
	m = make([][]int, h)
	f = make([][][4]int, h)
	for i := 0; i < h; i++ {
		m[i] = getInts(w)
		f[i] = make([][4]int, w)
		for j := 0; j < w; j++ {
			for k := 0; k < 4; k++ {
				f[i][j][k] = inf
			}
		}
	}

	bsf(0, 0)
	ans := inf
	for k := 0; k < 4; k++ {
		ans = min(ans, f[h-1][w-1][k])
	}
	if ans == inf {
		out(-1)
		return
	}
	out(ans)
}
