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

// 2つの線分に交点があるかどうかを判定
func judgeIentersected(ax, ay, bx, by, cx, cy, dx, dy int) bool {
	ta := (cx-dx)*(ay-cy) + (cy-dy)*(cx-ax)
	tb := (cx-dx)*(by-cy) + (cy-dy)*(cx-bx)
	tc := (ax-bx)*(cy-ay) + (ay-by)*(ax-cx)
	td := (ax-bx)*(dy-ay) + (ay-by)*(ax-dx)
	//	return tc*td < 0 && ta*tb < 0
	return tc*td <= 0 && ta*tb <= 0 // 端点を含む場合
}

func f(x0, y0, x1, y1 int) (int, int, int, int) {
	dx := x1 - x0
	dy := y1 - y0
	X0 := dx*100 + x0
	Y0 := dy*100 + y0
	X1 := -dx*100 + x0
	Y1 := -dy*100 + y0
	return X0, Y0, X1, Y1
}

func g(x0, y0, x1, y1 int) int {
	cnt := 0
	for i := 0; i < N; i++ {
		if judgeIentersected(x0, y0, x1, y1, n[i][0], n[i][1], n[i][2], n[i][3]) {
			cnt++
		}
	}
	return cnt
}

var N int
var n [][4]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	if N == 1 {
		out(1)
		return
	}
	n = make([][4]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < 4; j++ {
			n[i][j] = getI()
		}
	}
	// out(n)
	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if i == j {
				continue
			}
			x0, x1, y0, y1 := f(n[i][0], n[i][1], n[j][0], n[j][1])
			ans = max(ans, g(x0, x1, y0, y1))
			x0, x1, y0, y1 = f(n[i][2], n[i][3], n[j][2], n[j][3])
			ans = max(ans, g(x0, x1, y0, y1))
			x0, x1, y0, y1 = f(n[i][0], n[i][1], n[j][2], n[j][3])
			ans = max(ans, g(x0, x1, y0, y1))
			x0, x1, y0, y1 = f(n[i][2], n[i][3], n[j][0], n[j][1])
			ans = max(ans, g(x0, x1, y0, y1))
		}
	}
	out(ans)
}
