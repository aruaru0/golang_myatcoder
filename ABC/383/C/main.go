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

type pair struct {
	y, x int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, D := getI(), getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	const inf = int(1e9)
	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
		}
	}

	q := make([]pair, 0)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == 'H' {
				q = append(q, pair{i, j})
				dist[i][j] = 0
			}
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for len(q) != 0 {
		cur := q[0]
		cx, cy := cur.x, cur.y
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if s[py][px] == '#' {
				continue
			}
			if dist[py][px] > dist[cy][cx]+1 {
				dist[py][px] = dist[cy][cx] + 1
				q = append(q, pair{py, px})
			}
		}
	}
	cnt := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if dist[i][j] <= D {
				cnt++
			}
		}
	}
	out(cnt)
}
