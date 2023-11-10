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

func sign(x int) int {
	if x >= 0 {
		return 1
	}
	return -1
}

func isPointInsidePolygon(N int, x []int, y []int, a int, b int) bool {
	intersectCount := 0

	for i := 0; i < N; i++ {
		x1, y1 := x[i], y[i]
		x2, y2 := x[(i+1)%N], y[(i+1)%N]

		// 線分が水平でない場合
		if y1 != y2 {
			// 点が線分のy座標の範囲内にあるかチェック
			if (y1 <= b && y2 > b) || (y1 > b && y2 <= b) {
				// 線分とy座標が点のy座標と交差するかチェック
				if y2-y1 >= 0 {
					if (b-y1)*(x2-x1) > (a-x1)*(y2-y1) {
						intersectCount++
					}
				} else {
					if (b-y1)*(x2-x1) < (a-x1)*(y2-y1) {
						intersectCount++
					}
				}
			}
		}
	}

	// 交差点の数が奇数ならば点は多角形の内側にある
	return intersectCount%2 == 1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getI(), getI()
	}
	x = append(x, x[0])
	y = append(y, y[0])
	a, b := getI(), getI()

	if isPointInsidePolygon(N, x, y, a, b) {
		out("INSIDE")
	} else {
		out("OUTSIDE")
	}
}
