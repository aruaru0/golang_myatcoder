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

func big(lsx, sy, lgx, gy, b, k int) int {
	bsy := sy / b
	bgy := gy / b
	// 間にｘ方向の大通りが無い場合は、ｘはＫかけて移動、
	// 一度ｘに水平は通りに移動して移動（２通り）の３通りの最小
	if bsy == bgy {
		return min(abs(lsx-lgx)*k+abs(sy-gy), abs(lsx-lgx)+min(sy%b+gy%b, b-sy%b+b-gy%b))
	}
	// 間に水平な大通りがあれば、そこを通過すれば以下
	return abs(lsx-lgx) + abs(sy-gy)
}

func solve() {
	b, k, sx, sy, gx, gy := getI(), getI(), getI(), getI(), getI(), getI()
	// start周辺の大通り
	lsx := sx / b * b
	rsx := (sx + b) / b * b
	dsy := sy / b * b
	usy := (sy + b) / b * b
	// goal周辺の大通り
	lgx := gx / b * b
	rgx := (gx + b) / b * b
	dgy := gy / b * b
	ugy := (gy + b) / b * b

	// 大通りを使わない場合
	ans := (abs(sx-gx) + abs(sy-gy)) * k

	// 大通りを使う場合　(l,r,d,u)*(l,r,d,u)の4*4=16通列挙

	// lsx, sy -> lgx, gy
	ans = min(ans, (sx-lsx+gx-lgx)*k+big(lsx, sy, lgx, gy, b, k))

	// rsx, sy -> lgx, gy
	ans = min(ans, (abs(sx-rsx)+gx-lgx)*k+big(rsx, sy, lgx, gy, b, k))

	// lsx, sy -> rgx, gy
	ans = min(ans, (sx-lsx+abs(gx-rgx))*k+big(lsx, sy, rgx, gy, b, k))

	// rsx, sy -> rgx, gy
	ans = min(ans, (abs(sx-rsx)+abs(gx-rgx))*k+big(rsx, sy, rgx, gy, b, k))

	// sx, dsy -> gx, dgy
	ans = min(ans, (sy-dsy+gy-dgy)*k+big(dsy, sx, dgy, gx, b, k))

	// sx, dsy -> gx, ugy
	ans = min(ans, (sy-dsy+ugy-gy)*k+big(dsy, sx, ugy, gx, b, k))

	// sx, usy -> gx, dgy
	ans = min(ans, (usy-sy+gy-dgy)*k+big(usy, sx, dgy, gx, b, k))

	// sx, usy -> gx, ugy
	ans = min(ans, (usy-sy+ugy-gy)*k+big(usy, sx, ugy, gx, b, k))

	// lsx, sy -> gx, dgy
	ans = min(ans, (sx-lsx+gy-dgy)*k+abs(lsx-gx)+abs(sy-dgy))

	// rsx, sy -> gx, dgy
	ans = min(ans, (abs(sx-rsx)+gy-dgy)*k+abs(rsx-gx)+abs(sy-dgy))

	// lsx, sy -> gx, ugy
	ans = min(ans, (sx-lsx+ugy-gy)*k+abs(lsx-gx)+abs(sy-ugy))

	// rsx, sy -> gx, ugy
	ans = min(ans, (abs(sx-rsx)+ugy-gy)*k+abs(rsx-gx)+abs(sy-ugy))

	// sx, dsy -> lgx, gy
	ans = min(ans, (sy-dsy+gx-lgx)*k+abs(sx-lgx)+abs(dsy-gy))

	// sx, dsy -> rgx, gy
	ans = min(ans, (sy-dsy+rgx-gx)*k+abs(sx-rgx)+abs(dsy-gy))

	// sx, usy -> lgx, gy
	ans = min(ans, (usy-sy+gx-lgx)*k+abs(sx-lgx)+abs(usy-gy))

	// sx, usy -> rgx, gy
	ans = min(ans, (usy-sy+rgx-gx)*k+abs(sx-rgx)+abs(usy-gy))

	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	t := getI()
	for i := 0; i < t; i++ {
		solve()
	}
}
