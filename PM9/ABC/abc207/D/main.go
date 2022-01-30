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

func getAxis(n int) ([]float64, []float64) {
	x, y := make([]float64, n), make([]float64, n)
	gx, gy := 0.0, 0.0
	for i := 0; i < n; i++ {
		x[i], y[i] = getF(), getF()
		gx += x[i]
		gy += y[i]
	}
	gx /= float64(n)
	gy /= float64(n)
	for i := 0; i < n; i++ {
		x[i] -= gx
		y[i] -= gy
	}
	return x, y
}

const eps = 1e-6

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	sx, sy := getAxis(N)
	tx, ty := getAxis(N)

	// sx[0], sy[0] = 0, 0の場合の処理
	// 0,0の場合は回転角が求まらないので、処理が必要
	for i := 0; i < N; i++ {
		if sx[i] != 0 || sy[i] != 0 {
			sx[i], sx[0] = sx[0], sx[i]
			sy[i], sy[0] = sy[0], sy[i]
		}
	}

	// N=100なので、O(N^3)は間に合う
	for i := 0; i < N; i++ {
		// 角度を計算
		theta := math.Atan2(ty[i], tx[i]) - math.Atan2(sy[0], sx[0])
		// sを回転させてtになるか検査
		flag := true
		for j := 0; j < N; j++ {
			x := sx[j]*math.Cos(theta) - sy[j]*math.Sin(theta)
			y := sx[j]*math.Sin(theta) + sy[j]*math.Cos(theta)
			flag2 := false
			for k := 0; k < N; k++ {
				if math.Abs(x-tx[k]) < eps && math.Abs(y-ty[k]) < eps {
					flag2 = true
					// out("match", k, x, y, tx[k], ty[k])
					break
				}
			}
			// out("flag2 = ", flag2)
			if flag2 == false {
				flag = false
				break
			}
		}
		if flag {
			out("Yes")
			return
		}
	}
	out("No")
}
