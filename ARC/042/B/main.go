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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func getFloat() float64 {
	x := getInt()
	return float64(x)
}

type pair struct {
	x, y float64
}

func distance(x0, y0, x1, y1, x2, y2 float64) float64 {
	dist := math.Abs((y2-y1)*x0 - (x2-x1)*y0 + x2*y1 - y2*x1)
	dist /= math.Sqrt((y2-y1)*(y2-y1) + (x2-x1)*(x2-x1))
	return dist
}

func main() {
	sc.Split(bufio.ScanWords)
	x, y := getFloat(), getFloat()
	N := getInt()
	p := make([]pair, N+1)
	for i := 0; i < N; i++ {
		xx, yy := getFloat(), getFloat()
		p[i] = pair{xx, yy}
	}
	p[N] = p[0]
	ans := 1e20
	for i := 1; i <= N; i++ {
		dist := distance(x, y, p[i-1].x, p[i-1].y, p[i].x, p[i].y)
		// out(x, y, p[i-1].x, p[i-1].y, p[i].x, p[i].y)
		// out(dist)
		ans = math.Min(ans, dist)
	}
	out(ans)
}
