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

const inf = 100000

type pos struct {
	x, y int32
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H+2)
	a := make([][]int32, H+2)
	for i := 0; i < H+2; i++ {
		a[i] = make([]int32, W+2)
		for j := 0; j < W+2; j++ {
			a[i][j] = inf
		}
	}
	for i := 0; i < W+2; i++ {
		s[0] += "."
		s[H+1] += "."
	}
	for i := 0; i < H; i++ {
		s[i+1] = "." + getS() + "."
	}

	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	W += 2
	H += 2

	q := make([]pos, 0)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == '.' {
				ok := false
				for k := 0; k < 8; k++ {
					px := j + dx[k]
					py := i + dy[k]
					if px < 0 || px >= W || py < 0 || py >= H {
						continue
					}
					if s[py][px] != '.' {
						ok = true
					}
				}
				if ok {
					q = append(q, pos{int32(j), int32(i)})
				}
				a[i][j] = 0
			}
		}
	}

	for len(q) != 0 {
		e := q[0]
		q = q[1:]
		for i := 0; i < 8; i++ {
			px := int(e.x) + dx[i]
			py := int(e.y) + dy[i]
			if px < 0 || px >= W || py < 0 || py >= H {
				continue
			}
			if a[py][px] > a[e.y][e.x]+1 {
				a[py][px] = a[e.y][e.x] + 1
				q = append(q, pos{int32(px), int32(py)})
			}
		}
	}

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			ans = max(ans, int(a[i][j]))
		}
	}
	// for i := 0; i < H; i++ {
	// 	out(i, a[i])
	// }
	out(ans)
}
