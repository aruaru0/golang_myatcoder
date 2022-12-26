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

const size = 420
const offset = 210
const inf int = 1e18

type pos struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, X, Y := getI(), getI()+offset, getI()+offset

	var a [size][size]bool
	var d [size][size]int

	for i := 0; i < N; i++ {
		x, y := getI()+offset, getI()+offset
		a[y][x] = true
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			d[i][j] = inf
		}
	}

	q := []pos{}
	q = append(q, pos{offset, offset})
	d[offset][offset] = 0

	dx := []int{1, 0, -1, 1, -1, 0}
	dy := []int{1, 1, 1, 0, 0, -1}

	for len(q) != 0 {
		cur := q[0]
		cx, cy := cur.x, cur.y
		q = q[1:]
		for i := 0; i < 6; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px < 0 || px >= size || py < 0 || py >= size {
				continue
			}
			if a[py][px] {
				continue
			}
			if d[py][px] > d[cy][cx]+1 {
				d[py][px] = d[cy][cx] + 1
				q = append(q, pos{px, py})
			}
		}
	}

	// for i := offset; i < offset+10; i++ {
	// 	out(d[i][offset : offset+10])
	// }
	if d[Y][X] == inf {
		out(-1)
		return
	}
	out(d[Y][X])

}
