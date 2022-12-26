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

var N, M int
var s [][]byte

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

type pos struct {
	y, x int
}

func check(y, x int) int {
	t := make([][]byte, N)
	for i := 0; i < N; i++ {
		t[i] = make([]byte, M)
	}
	q := make([]pos, 0)
	q = append(q, pos{y, x})
	t[y][x] = 1
	cnt := 1
	for len(q) != 0 {
		cx := q[0].x
		cy := q[0].y
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px < 0 || py < 0 || px >= M || py >= N {
				continue
			}
			if s[py][px] == '#' {
				continue
			}
			if t[py][px] == 0 {
				t[py][px] = 1
				q = append(q, pos{py, px})
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	s = make([][]byte, N)
	tot := 0
	for i := 0; i < N; i++ {
		s[i] = []byte(getS())
		for j := 0; j < M; j++ {
			if s[i][j] == '.' {
				tot++
			}
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if s[i][j] == '#' {
				s[i][j] = '.'
				cnt := check(i, j)
				if cnt == tot+1 {
					ans++
				}
				s[i][j] = '#'
			}
		}
	}
	out(ans)
}
