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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func norm(up, dn int, slope bool) (int, int) {
	// 分母が負の場合は正にする
	if dn < 0 {
		up *= -1
		dn *= -1
	}
	// 分母が０の場合はSlope＝trueならupを１にする
	// ※水平の場合の特殊処理
	if dn == 0 {
		if slope {
			up = 1
		}
	} else {
		// gcdで割ってdx,dyを正規化
		g := gcd(abs(up), dn)
		up /= g
		dn /= g
	}
	return up, dn
}

var N, K int
var X, Y []int

type P struct {
	dy, dx, bu, bd int
}

func getParam(a, b int) P {
	// dx, dyを求める
	dx := X[a] - X[b]
	dy := Y[a] - Y[b]
	dy, dx = norm(dy, dx, true)

	// 切片位置を求める
	// y = dy/dx * x + b
	// b = y - dy/dx * x = (dx*y - dy*x)/dx = bu/bd
	bu := Y[a]*dx - dy*X[a]
	bd := dx
	bu, bd = norm(bu, bd, false)

	return P{dy, dx, bu, bd}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K = getI(), getI()
	if K == 1 {
		out("Infinity")
		return
	}
	X = make([]int, N)
	Y = make([]int, N)
	for i := 0; i < N; i++ {
		X[i], Y[i] = getI(), getI()
	}

	cnt := make(map[P]int)
	for a := 0; a < N; a++ {
		for b := a + 1; b < N; b++ {
			cnt[getParam(a, b)]++
		}
	}

	cmb := make(map[int]int)
	for n := K; n < N+1; n++ {
		cmb[n*(n-1)/2] = 1
	}

	ans := 0
	for _, e := range cnt {
		ans += cmb[e]
	}
	out(ans)
}
