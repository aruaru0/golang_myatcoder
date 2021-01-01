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

var H, W, K int
var c []string

func check(h, w int) int {
	b := make([][]int, H)
	for i := 0; i < H; i++ {
		b[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		if (h>>i)&1 == 1 {
			for j := 0; j < W; j++ {
				b[i][j] = 1
			}
		}
	}

	for i := 0; i < W; i++ {
		if (w>>i)&1 == 1 {
			for j := 0; j < H; j++ {
				b[j][i] = 1
			}
		}
	}

	tot := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if b[i][j] == 0 && c[i][j] == '#' {
				tot++
			}
		}
	}
	return tot
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, K = getI(), getI(), getI()
	c = make([]string, H)
	for i := 0; i < H; i++ {
		c[i] = getS()
	}

	h := 1 << H
	w := 1 << W
	cnt := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ret := check(i, j)
			if ret == K {
				cnt++
			}
		}
	}
	out(cnt)
}
