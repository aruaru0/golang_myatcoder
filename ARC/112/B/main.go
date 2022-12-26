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

type pair struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	b, c := getI(), getI()

	//　最初に-1倍、最後に-1倍する組み合わせで
	// 到達できる範囲をa(area)に入れる
	a := make([]pair, 0)
	// - -
	a = append(a, pair{b, b - c/2})
	// x-1 -
	a = append(a, pair{-b, -b - (c-1)/2})
	// - x-1
	a = append(a, pair{-(b - (c-1)/2), -b})
	// x-1 x-1
	if c >= 2 {
		a = append(a, pair{-(-b - (c-2)/2), b})
	}

	// imos法のデータに変換
	x := make([]pair, 0)
	for _, e := range a {
		f, t := e.x, e.y
		if f > t {
			f, t = t, f
		}
		x = append(x, pair{f, +1})
		x = append(x, pair{t + 1, -1})
	}

	// ソートして並べる
	sort.Slice(x, func(i, j int) bool {
		return x[i].x < x[j].x
	})

	// 値の存在する範囲の数を調べる
	prev := -int(1e18 + 100)
	tot := 0
	ans := 0
	for _, e := range x {
		if tot > 0 {
			ans += e.x - prev
		}
		prev = e.x
		tot += e.y
	}
	out(ans)
}
