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

var H, W int
var ch, cw int
var dh, dw int
var S []string
var u [][]int

const inf = int(1e9)

type next struct {
	x, y int
}

func bsf(s []next, c int) []next {
	q := s
	nq := make([]next, 0)
	for len(q) != 0 {
		cx := q[0].x
		cy := q[0].y
		q = q[1:]
		for dy := -2; dy <= 2; dy++ {
			for dx := -2; dx <= 2; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				x := cx + dx
				y := cy + dy
				if x < 0 || x >= W || y < 0 || y >= H {
					continue
				}
				if S[y][x] == '#' {
					continue
				}
				if u[y][x] > c {
					if dx == 0 && abs(dy) == 1 || dy == 0 && abs(dx) == 1 {
						u[y][x] = c
						q = append(q, next{x, y})
					} else {
						u[y][x] = c + 1
						nq = append(nq, next{x, y})
					}
				}
			}
		}
	}
	return nq
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W = getInt(), getInt()
	ch, cw = getInt()-1, getInt()-1
	dh, dw = getInt()-1, getInt()-1
	S = make([]string, H)
	u = make([][]int, H)
	for i := 0; i < H; i++ {
		S[i] = getString()
		u[i] = make([]int, W)
		for j := 0; j < W; j++ {
			u[i][j] = inf
		}
	}

	u[ch][cw] = 0
	n := make([]next, 0)
	n = append(n, next{cw, ch})
	for i := 0; len(n) != 0; i++ {
		n = bsf(n, i)
		// for i := 0; i < H; i++ {
		// 	out(u[i])
		// }
		// out("----------------")
	}
	if u[dh][dw] == inf {
		out(-1)
		return
	}
	out(u[dh][dw])
}
