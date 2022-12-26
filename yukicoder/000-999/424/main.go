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

var h, w, sx, sy, gx, gy int
var b [][]int
var d [][]int

type pos struct {
	x, y int
}

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, -1, 1}

func bfs(sx, sy int) {
	d[sy][sx] = 0
	q := make([]pos, 0)
	q = append(q, pos{sx, sy})
	for len(q) != 0 {
		cx := q[0].x
		cy := q[0].y
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px < 0 || py < 0 || px >= w || py >= h {
				continue
			}
			if abs(b[py][px]-b[cy][cx]) <= 1 {
				if d[py][px] == inf {
					d[py][px] = d[cy][cx] + 1
					q = append(q, pos{px, py})
				}
			}
		}
		for i := 0; i < 4; i++ {
			px := cx + dx[i]*2
			py := cy + dy[i]*2
			if px < 0 || py < 0 || px >= w || py >= h {
				continue
			}
			if b[py][px] == b[cy][cx] && b[(cy+py)/2][(cx+px)/2] < b[py][px] {
				if d[py][px] == inf {
					d[py][px] = d[cy][cx] + 1
					q = append(q, pos{px, py})
				}
			}
		}
	}
}

const inf = int(1e9)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	h, w = getI(), getI()
	sy, sx, gy, gx = getI()-1, getI()-1, getI()-1, getI()-1
	b = make([][]int, h)
	d = make([][]int, h)
	for i := 0; i < h; i++ {
		b[i] = make([]int, w)
		d[i] = make([]int, w)
		s := getS()
		for j := 0; j < w; j++ {
			b[i][j] = int(s[j] - '0')
			d[i][j] = inf
		}
	}

	bfs(sx, sy)
	// for i := 0; i < h; i++ {
	// 	out(b[i])
	// }
	// out("----")
	// for i := 0; i < h; i++ {
	// 	out(d[i])
	// }
	if d[gy][gx] != inf {
		out("YES")
		return
	}
	out("NO")
}
