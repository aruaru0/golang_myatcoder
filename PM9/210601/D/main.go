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
	H, W, _ := getI(), getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	ans := make([][]int, H)
	for y := 0; y < H; y++ {
		ans[y] = make([]int, W)
	}
	cnt := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '#' {
				cnt++
				ans[y][x] = cnt
			}
		}
	}

	for y := 0; y < H; y++ {
		for x := 1; x < W; x++ {
			if ans[y][x] == 0 {
				ans[y][x] = ans[y][x-1]
			}
		}
		for x := W - 1; x > 0; x-- {
			if ans[y][x-1] == 0 {
				ans[y][x-1] = ans[y][x]
			}
		}
	}
	for x := 0; x < W; x++ {
		for y := 1; y < H; y++ {
			if ans[y][x] == 0 {
				ans[y][x] = ans[y-1][x]
			}
		}
		for y := H - 1; y > 0; y-- {
			if ans[y-1][x] == 0 {
				ans[y-1][x] = ans[y][x]
			}
		}
	}

	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			fmt.Fprint(wr, ans[y][x], " ")
		}
		out()
	}
}
