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
	a := make([][]int, H+2)
	for i := 0; i < H+2; i++ {
		a[i] = make([]int, W+2)
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

	W += 2
	H += 2

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '.' {
				a[y][x] = 0
			}
		}
	}

	ans := 0
	for y := 1; y < H; y++ {
		for x := 1; x < W; x++ {
			if s[y][x] == '#' {
				a[y][x] = nmin(a[y-1][x-1], a[y-1][x], a[y][x-1]) + 1
				ans = max(ans, a[y][x])
			}
		}
	}
	out((ans + 1) / 2)
}