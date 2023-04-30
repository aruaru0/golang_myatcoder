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

var used [][]bool
var c []string
var H, W int

func dfs(h, w int) int {
	used[h][w] = true
	dx := []int{-1, 1, -1, 1}
	dy := []int{-1, -1, 1, 1}
	ret := 1
	for i := 0; i < 4; i++ {
		py := h + dy[i]
		px := w + dx[i]
		if py < 0 || py >= H || px < 0 || px >= W {
			continue
		}
		if used[py][px] == true {
			continue
		}
		if c[py][px] != '#' {
			continue
		}
		ret += dfs(py, px)
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	c = make([]string, H)
	for i := 0; i < H; i++ {
		c[i] = getS()
	}
	used = make([][]bool, H)
	for i := 0; i < H; i++ {
		used[i] = make([]bool, W)
	}

	n := make([]int, min(H, W)+1)
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if used[h][w] == true {
				continue
			}
			if c[h][w] != '#' {
				continue
			}
			ret := dfs(h, w)

			n[ret/4]++
		}
	}
	for i := 1; i < len(n); i++ {
		fmt.Fprint(wr, n[i], " ")
	}
	out()
}
