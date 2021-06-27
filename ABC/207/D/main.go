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

type pos struct {
	x, y float64
}

func f(n int) ([]float64, []float64) {
	x := make([]float64, n)
	y := make([]float64, n)
	gx, gy := 0.0, 0.0
	for i := 0; i < n; i++ {
		a, b := getF(), getF()
		gx += a
		gy += b
		x[i] = a
		y[i] = b
	}
	gx /= float64(n)
	gy /= float64(n)
	for i := 0; i < n; i++ {
		x[i] -= gx
		y[i] -= gy
	}
	return x, y
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	// この手の考えるのは簡単だけど、浮動小数点の誤差だけの問題は
	// あんまり面白くない。多分、作る方はよくても解く方はうんざりする問題

	a, b := f(N)
	c, d := f(N)

	// refx, refy = 0,0を避ける（中心を避ける）
	refx, refy := a[0], b[0]
	for i := 0; i < N; i++ {
		if a[i] != 0 || b[i] != 0 {
			refx, refy = a[i], b[i]
			break
		}
	}

	const eps = 1e-6
	for i := 0; i < N; i++ {
		// どれかと一致する角度を計算する
		angle := math.Atan2(d[i], c[i]) - math.Atan2(refy, refx)
		flag := true
		for j := 0; j < N; j++ {
			// 角度を回転させて
			A := a[j]*math.Cos(angle) - b[j]*math.Sin(angle)
			B := a[j]*math.Sin(angle) + b[j]*math.Cos(angle)
			flag2 := false
			// 一致するやつがあるか確認
			for k := 0; k < N; k++ {
				if math.Abs(A-c[k]) <= eps && math.Abs(B-d[k]) <= eps {
					flag2 = true
				}
			}
			flag = flag && flag2
		}
		if flag {
			out("Yes")
			return
		}
	}
	out("No")
}
