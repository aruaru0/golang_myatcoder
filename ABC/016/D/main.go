package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

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

// 2つの線分に交点があるかどうかを判定
func judgeIentersected(ax, ay, bx, by, cx, cy, dx, dy int) bool {
	ta := (cx-dx)*(ay-cy) + (cy-dy)*(cx-ax)
	tb := (cx-dx)*(by-cy) + (cy-dy)*(cx-bx)
	tc := (ax-bx)*(cy-ay) + (ay-by)*(ax-cx)
	td := (ax-bx)*(dy-ay) + (ay-by)*(ax-dx)

	return tc*td < 0 && ta*tb < 0
	// return tc * td <= 0 && ta * tb <= 0; // 端点を含む場合
}

func main() {
	sc.Split(bufio.ScanWords)
	ax, ay, bx, by := getInt(), getInt(), getInt(), getInt()
	N := getInt()
	x := make([]int, N+1)
	y := make([]int, N+1)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
	}
	x[N], y[N] = x[0], y[0]
	cnt := 0
	for i := 1; i <= N; i++ {
		ret := judgeIentersected(ax, ay, bx, by, x[i], y[i], x[i-1], y[i-1])
		if ret {
			cnt++
		}
	}
	out(cnt/2 + 1)
}
