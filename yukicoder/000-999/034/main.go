package main

import (
	"bufio"
	"fmt"
	"math"
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

var N, V, sx, sy, gx, gy int
var m [][]int
var dist [][]int
var cnt [][]int

const inf = math.MaxInt64 / 16

type pair struct {
	x, y int
}

func bsf(sx, sy int) {
	dist = make([][]int, N)
	cnt = make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		cnt[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = inf
			cnt[i][j] = inf
		}
	}
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	dist[sy][sx] = 0
	cnt[sy][sx] = 0
	q := make([]pair, 0)
	q = append(q, pair{sx, sy})
	for len(q) != 0 {
		p := q[0]
		if p.x == gx && p.y == gy {
			break
		}
		q = q[1:]
		for i := 0; i < 4; i++ {
			x := p.x + dx[i]
			y := p.y + dy[i]
			if x < 0 || x >= N || y < 0 || y >= N {
				continue
			}
			if cnt[p.y][p.x]+m[y][x] >= V {
				continue
			}
			if cnt[y][x] > cnt[p.y][p.x]+m[y][x] {
				dist[y][x] = dist[p.y][p.x] + 1
				cnt[y][x] = cnt[p.y][p.x] + m[y][x]
				q = append(q, pair{x, y})
			}
		}
	}

}

func main() {
	sc.Split(bufio.ScanWords)
	N, V, sx, sy, gx, gy = getInt(), getInt(),
		getInt()-1, getInt()-1, getInt()-1, getInt()-1
	m = make([][]int, N)
	for i := 0; i < N; i++ {
		m[i] = getInts(N)
	}
	bsf(sx, sy)
	// for i := 0; i < N; i++ {
	// 	out(cnt[i])
	// }

	if dist[gy][gx] != inf {
		out(dist[gy][gx])
		return
	}
	out(-1)
}
