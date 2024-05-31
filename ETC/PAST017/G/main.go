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

var H, W, N int
var g []string
var s string
var used [][]bool

func check(idx, h, w int) bool {
	if idx == N-1 {
		return true
	}
	used[h][w] = true

	for dh := -1; dh <= 1; dh++ {
		for dw := -1; dw <= 1; dw++ {
			if dh == 0 && dw == 0 {
				continue
			}
			ph, pw := h+dh, w+dw
			if ph < 0 || ph >= H || pw < 0 || pw >= W {
				continue
			}
			if used[ph][pw] {
				continue
			}
			if g[ph][pw] != s[idx+1] {
				continue
			}
			// out(ph, pw, string(g[ph][pw]), string(s[idx+1]), "->check")
			ret := check(idx+1, ph, pw)
			if ret {
				return true
			}
		}
	}
	used[h][w] = false
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	g = make([]string, H)
	for i := 0; i < H; i++ {
		g[i] = getS()
	}
	N = getI()
	s = getS()
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if g[h][w] == s[0] {
				used = make([][]bool, H)
				for i := 0; i < H; i++ {
					used[i] = make([]bool, W)
				}
				ret := check(0, h, w)
				if ret {
					out("Yes")
					return
				}
			}
		}
	}
	out("No")
}
