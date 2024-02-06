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
	x1, y1, x2, y2 int
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
	}

	x1, x2, y1, y2 := -1, -1, -1, -1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s[i][j] == 'P' {
				if x1 == -1 {
					x1, y1 = i, j
				} else {
					x2, y2 = i, j
				}
			}
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	q := []pos{}
	var d [60][60][60][60]int
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			for k := 0; k < 60; k++ {
				for l := 0; l < 60; l++ {
					d[i][j][k][l] = inf
				}
			}
		}
	}

	d[x1][y1][x2][y2] = 0

	q = append(q, pos{x1, y1, x2, y2})
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			px1, py1 := cur.x1+dx[i], cur.y1+dy[i]
			if px1 < 0 || px1 >= N || py1 < 0 || py1 >= N || s[px1][py1] == '#' {
				px1, py1 = cur.x1, cur.y1
			}
			px2, py2 := cur.x2+dx[i], cur.y2+dy[i]
			if px2 < 0 || px2 >= N || py2 < 0 || py2 >= N || s[px2][py2] == '#' {
				px2, py2 = cur.x2, cur.y2
			}
			if d[px1][py1][px2][py2] == inf {
				// out(cur, "->", px1, py1, px2, py2)
				d[px1][py1][px2][py2] = d[cur.x1][cur.y1][cur.x2][cur.y2] + 1
				q = append(q, pos{px1, py1, px2, py2})
			}
		}
	}

	ans := inf
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans = min(ans, d[i][j][i][j])
		}
	}

	if ans == inf {
		out("-1")
	} else {
		out(ans)
	}

}
