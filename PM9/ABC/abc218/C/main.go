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

// rotate 90 degree clockwise
func rotate90(s [][]byte) [][]byte {
	N := len(s)
	ret := make([][]byte, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]byte, N)
		for j := 0; j < N; j++ {
			ret[i][j] = s[N-1-j][i]
		}
	}
	return ret
}

func topleft(s [][]byte) (int, int) {
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if s[y][x] == '#' {
				return x, y
			}
		}
	}
	return 0, 0
}

func bottomleft(s [][]byte) (int, int) {
	for y := N - 1; y >= 0; y-- {
		for x := 0; x < N; x++ {
			if s[y][x] == '#' {
				return x, y
			}
		}
	}
	return 0, 0
}

func match(sx, sy, tx, ty int, s, t [][]byte) bool {
	offsetX := tx - sx
	offsetY := ty - sy
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			px := x + offsetX
			py := y + offsetY
			if 0 <= px && px < N && 0 <= py && py < N {
				if s[y][x] != t[py][px] {
					return false
				}
			} else {
				if s[y][x] == '#' {
					return false
				}
			}
		}
	}
	return true
}

var N int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	s := make([][]byte, N)
	t := make([][]byte, N)
	for i := 0; i < N; i++ {
		s[i] = []byte(getS())
	}
	for i := 0; i < N; i++ {
		t[i] = []byte(getS())
	}

	scnt, tcnt := 0, 0
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if s[y][x] == '#' {
				scnt++
			}
			if t[y][x] == '#' {
				tcnt++
			}
		}
	}
	if scnt != tcnt {
		out("No")
		return
	}

	tx, ty := topleft(t)
	for i := 0; i < 4; i++ {
		sx, sy := topleft(s)
		ok := match(sx, sy, tx, ty, s, t)
		if ok {
			out("Yes")
			return
		}
		s = rotate90(s)
	}
	out("No")
}
