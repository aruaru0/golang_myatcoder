package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	sc.Split(bufio.ScanWords)
	A, B, H, M := float64(getInt()),
		float64(getInt()),
		float64(getInt()),
		float64(getInt())

	H = H + M/60
	// out(math.Pi * H / 6)
	H = (H * math.Pi) / 6
	x0 := A * math.Cos(H)
	y0 := A * math.Sin(H)

	M = M * math.Pi / 30
	x1 := B * math.Cos(M)
	y1 := B * math.Sin(M)
	// out(math.Sqrt(x0*x0 + y0*y0))
	// out(math.Sqrt(x1*x1 + y1*y1))

	dx := x0 - x1
	dy := y0 - y1
	l := math.Sqrt(dx*dx + dy*dy)
	// out(x0, y0, x1, y1)
	out(l)
}
