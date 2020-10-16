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
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	R, K := getI(), getI()
	H, W := getI(), getI()
	c := make([][]byte, H)

	for i := 0; i < H; i++ {
		c[i] = []byte(getS())
	}

	var d [][]byte
	var h, w int
	switch R {
	case 0:
		h, w = H, W
		d = make([][]byte, H)
		for y := 0; y < H; y++ {
			d[y] = make([]byte, W)
		}
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				d[y][x] = c[y][x]
			}
		}
	case 90:
		h, w = W, H
		d = make([][]byte, W)
		for y := 0; y < W; y++ {
			d[y] = make([]byte, H)
		}
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				d[x][H-1-y] = c[y][x]
			}
		}
	case 180:
		h, w = H, W
		d = make([][]byte, H)
		for y := 0; y < H; y++ {
			d[y] = make([]byte, W)
		}
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				d[H-1-y][W-1-x] = c[y][x]
			}
		}
	case 270:
		h, w = W, H
		d = make([][]byte, W)
		for y := 0; y < W; y++ {
			d[y] = make([]byte, H)
		}
		for y := 0; y < H; y++ {
			for x := 0; x < W; x++ {
				d[W-1-x][y] = c[y][x]
			}
		}
	}

	for y := 0; y < h*K; y++ {
		yy := y / K
		for x := 0; x < w; x++ {
			for i := 0; i < K; i++ {
				fmt.Fprint(wr, string(d[yy][x]))
			}
		}
		out()
	}
}
