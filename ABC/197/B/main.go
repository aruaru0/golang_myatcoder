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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, Y, X := getI(), getI(), getI()-1, getI()-1
	s := make([]string, H)
	h := make([][]int, H)
	w := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
		h[i] = make([]int, W)
		w[i] = make([]int, W)
	}

	for y := 0; y < H; y++ {
		cnt := 0
		for x := 0; x < W; x++ {
			if s[y][x] == '#' {
				cnt = 0
				continue
			}
			cnt++
			w[y][x] = cnt
		}
	}
	for y := 0; y < H; y++ {
		cnt := w[y][W-1]
		for x := W - 1; x >= 0; x-- {
			if s[y][x] == '#' {
				cnt = 0
				continue
			}
			cnt = max(cnt, w[y][x])
			w[y][x] = cnt
		}
	}

	for x := 0; x < W; x++ {
		cnt := 0
		for y := 0; y < H; y++ {
			if s[y][x] == '#' {
				cnt = 0
				continue
			}
			cnt++
			h[y][x] = cnt
		}
	}
	for x := 0; x < W; x++ {
		cnt := h[H-1][x]
		for y := H - 1; y >= 0; y-- {
			if s[y][x] == '#' {
				cnt = 0
				continue
			}
			cnt = max(cnt, h[y][x])
			h[y][x] = cnt
		}
	}

	out(h[Y][X] + w[Y][X] - 1)
}
