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
	N := getI()
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	dir := 0
	cnt := 0
	x, y := 0, 0

	for N == 1 {
		out("001")
		return
	}

	for cnt != N*N {
		cnt++
		a[y][x] = cnt
		nx, ny := x+dx[dir], y+dy[dir]
		if nx < 0 || ny < 0 || nx >= N || ny >= N {
			dir = (dir + 1) % 4
			nx, ny = x+dx[dir], y+dy[dir]
		}
		if a[ny][nx] != 0 {
			dir = (dir + 1) % 4
			nx, ny = x+dx[dir], y+dy[dir]
		}
		x, y = nx, ny
	}

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			fmt.Fprintf(wr, "%3.3d ", a[y][x])
		}
		out()
	}
}
