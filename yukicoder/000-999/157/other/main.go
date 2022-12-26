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

var W, H int
var s []string
var n [][]int

type pair struct {
	x, y int
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func bsf(sx, sy, c int) int {
	if s[sy][sx] == '#' || n[sy][sx] != 0 {
		return c
	}
	q := make([]pair, 0)
	q = append(q, pair{sx, sy})
	n[sy][sx] = c
	for len(q) != 0 {
		cx := q[0].x
		cy := q[0].y
		q = q[1:]
		for i := 0; i < 4; i++ {
			x := cx + dx[i]
			y := cy + dy[i]
			if x < 0 || x >= W || y < 0 || y >= H {
				continue
			}
			if s[y][x] == '#' {
				continue
			}
			if n[y][x] == 0 {
				n[y][x] = c
				q = append(q, pair{x, y})
			}
		}
	}
	return c + 1
}

func main() {
	sc.Split(bufio.ScanWords)
	W, H = getInt(), getInt()
	s = make([]string, H)
	n = make([][]int, H)
	m := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
		n[i] = make([]int, W)
		m[i] = make([]int, W)
	}

	cnt := 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			cnt = bsf(j, i, cnt)
		}
	}

	for cnt := 0; ; cnt++ {
		for i := 0; i < H; i++ {
			copy(m[i], n[i])
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if n[i][j] == 1 {
					for k := 0; k < 4; k++ {
						x := j + dx[k]
						y := i + dy[k]
						if x < 0 || x >= W || y < 0 || y >= H {
							continue
						}
						if n[y][x] == 2 {
							out(cnt)
							return
						} else {
							m[y][x] = 1
						}
					}
				}
			}
		}
		m, n = n, m
		// for i := 0; i < H; i++ {
		// 	out(n[i])
		// }
		// out("--------")
	}
}
