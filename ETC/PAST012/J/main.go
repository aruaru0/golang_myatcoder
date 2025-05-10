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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func rotate(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			tmp := grid[i][j]
			grid[i][j] = grid[n-1-j][i]
			grid[n-1-j][i] = grid[n-1-i][n-1-j]
			grid[n-1-i][n-1-j] = grid[j][n-1-i]
			grid[j][n-1-i] = tmp
		}
	}
	return grid
}

func flipV(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n/2; i++ {
		nr := n - i - 1
		grid[nr], grid[i] = grid[i], grid[nr]
	}
	return grid
}

func flipH(grid [][]int) [][]int {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			nc := n - j - 1
			grid[i][nc], grid[i][j] = grid[i][j], grid[i][nc]
		}
	}
	return grid
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, q := getI(), getI()

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	rotateCnt := 0
	isFlipV, isFlipH := false, false

	for i := 0; i < q; i++ {
		t := getI()
		switch t {
		case 1:
			y, x := getI()-1, getI()-1
			for i := 0; i < 4-rotateCnt; i++ {
				y, x = x, n-1-y
			}
			if isFlipV {
				y = n - y - 1
			}
			if isFlipH {
				x = n - x - 1
			}
			grid[y][x] = 1 - grid[y][x]
		case 2:
			c := getS()
			if c == "A" {
				rotateCnt++
			} else {
				rotateCnt--
			}
			rotateCnt = (rotateCnt + 4) % 4
		case 3:
			c := getS()
			if c == "A" {
				if rotateCnt%2 == 0 {
					isFlipV = !isFlipV
				} else {
					isFlipH = !isFlipH
				}
			} else {
				if rotateCnt%2 == 0 {
					isFlipH = !isFlipH
				} else {
					isFlipV = !isFlipV
				}
			}
		}
	}

	if isFlipV {
		grid = flipV(grid)
	}
	if isFlipH {
		grid = flipH(grid)
	}
	for i := 0; i < rotateCnt; i++ {
		grid = rotate(grid)
	}

	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}
