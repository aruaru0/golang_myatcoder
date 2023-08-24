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

func dist(ax, ay, bx, by, cx, cy float64) float64 {
	// ベクトルABおよびACを計算
	ABx := bx - ax
	ABy := by - ay
	ACx := cx - ax
	ACy := cy - ay

	// ベクトルABとACの内積を計算
	dotProduct := ABx*ACx + ABy*ACy

	// 点Cが点Aから点Bに向かうベクトルABの途中にある場合
	if dotProduct > 0 {
		ABMagnitude := math.Sqrt(ABx*ABx + ABy*ABy)
		return math.Abs(ABx*ACy-ABy*ACx) / ABMagnitude
	}

	// 点Cが点Bから点Aに向かうベクトルBAの途中にある場合
	BCx := cx - bx
	BCy := cy - by
	BCMagnitude := math.Sqrt(BCx*BCx + BCy*BCy)
	return math.Abs(BCx*ACy-BCy*ACx) / BCMagnitude
}

func dist2(x0, y0, x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	a2 := a * a
	b2 := b * b
	r2 := a2 + b2
	tt := -(a*(x1-x0) + b*(y1-y0))
	if tt < 0 {
		return math.Sqrt((x1-x0)*(x1-x0) + (y1-y0)*(y1-y0))
	}
	if tt > r2 {
		return math.Sqrt((x2-x0)*(x2-x0) + (y2-y0)*(y2-y0))
	}
	var f1 = a*(y1-y0) - b*(x1-x0)
	return math.Sqrt((f1 * f1) / r2)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	ax, ay := getF(), getF()
	bx, by := getF(), getF()
	cx, cy := getF(), getF()

	distance := dist2(ax, ay, bx, by, cx, cy)
	out(distance)
}
