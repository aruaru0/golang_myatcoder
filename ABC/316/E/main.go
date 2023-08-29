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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([][]byte, H)
	for i := 0; i < H; i++ {
		a[i] = []byte(getS())
	}
	sx, sy := 0, 0
	gx, gy := 0, 0
	dist := make([][]int, H)
	for h := 0; h < H; h++ {
		dist[h] = make([]int, W)
		for w := 0; w < W; w++ {
			if a[h][w] == 'S' {
				sx, sy = w, h
			}
			if a[h][w] == 'G' {
				gx, gy = w, h
			}
			dist[h][w] = inf
		}
	}

	// // left
	for h := 0; h < H; h++ {
		ok := false
		for w := 0; w < W; w++ {
			if a[h][w] == '>' {
				ok = true
			} else if a[h][w] != '.' && a[h][w] != '!' && a[h][w] != 'G' && a[h][w] != 'S' {
				ok = false
			}
			if ok && a[h][w] != '>' {
				a[h][w] = '!'
			}
		}
	}

	// // right
	for h := 0; h < H; h++ {
		ok := false
		for w := W - 1; w >= 0; w-- {
			if a[h][w] == '<' {
				ok = true
			} else if a[h][w] != '.' && a[h][w] != '!' && a[h][w] != 'G' && a[h][w] != 'S' {
				ok = false
			}
			if ok && a[h][w] != '<' {
				a[h][w] = '!'
			}
		}
	}

	// down
	for w := 0; w < W; w++ {
		ok := false
		for h := 0; h < H; h++ {
			if a[h][w] == 'v' {
				ok = true
			} else if a[h][w] != '.' && a[h][w] != '!' && a[h][w] != 'G' && a[h][w] != 'S' {
				ok = false
			}
			if ok && a[h][w] != 'v' {
				a[h][w] = '!'
			}
		}
	}

	// up
	for w := 0; w < W; w++ {
		ok := false
		for h := H - 1; h >= 0; h-- {
			if a[h][w] == '^' {
				ok = true
			} else if a[h][w] != '.' && a[h][w] != '!' && a[h][w] != 'G' && a[h][w] != 'S' {
				ok = false
			}
			if ok && a[h][w] != '^' {
				a[h][w] = '!'
			}
		}
	}

	// for h := 0; h < H; h++ {
	// 	out(string(a[h]))
	// }

	// out(sx, sy, "-", gx, gy)

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	q := []pos{pos{sx, sy}}
	dist[sy][sx] = 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cur.x + dx[i]
			py := cur.y + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if a[py][px] != '.' && a[py][px] != 'G' && a[py][px] != 'S' {
				continue
			}
			if dist[py][px] > dist[cur.y][cur.x]+1 {
				dist[py][px] = dist[cur.y][cur.x] + 1
				q = append(q, pos{px, py})
			}
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }

	if dist[gy][gx] == inf {
		out(-1)
		return
	}
	out(dist[gy][gx])
}
