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

func calc(sx, sy int) float64 {
	// out(sx, sy)
	cnt := 0.0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if p[y][x] == '1' {
				dx := (x + 1 - sx)
				dy := (y + 1 - sy)
				cnt += math.Sqrt(float64(dx*dx + dy*dy))
			}
		}
	}
	return cnt
}

var H, W int
var p []string

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	p = make([]string, H)
	for i := 0; i < H; i++ {
		p[i] = getS()
	}

	ans := math.MaxFloat64
	for y := 1; y <= H; y++ {
		d := calc(0, y)
		ans = math.Min(ans, d)
	}
	for y := 1; y <= H; y++ {
		d := calc(W+1, y)
		ans = math.Min(ans, d)
	}
	for x := 1; x <= W; x++ {
		d := calc(x, 0)
		ans = math.Min(ans, d)
	}
	for x := 1; x <= W; x++ {
		d := calc(x, H+1)
		ans = math.Min(ans, d)
	}
	out(ans)
}
