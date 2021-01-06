package main

import (
	"bufio"
	"fmt"
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

var memo [101][2][101]float64
var saved [101][2][101]bool

// tgl : 前回必勝法を使った(1)・使わなかった(0)
func rec(n, tgl, p int) float64 {
	if n == N {
		return 1
	}
	if saved[n][tgl][p] == true {
		return memo[n][tgl][p]
	}
	np := p
	if n != 0 { // 最初以外は確率を更新
		if tgl == 1 {
			np = max(0, p-q)
		} else {
			np = min(100, p+q)
		}
	}
	fp := float64(np) / 100.0
	ret := 0.0
	// 必勝法を使って勝った場合
	ret = fp / 2.0
	// 必勝法を使わずに勝った場合
	ret += (1 - fp) / 3.0
	// 必勝法を使って引き分け
	ret += fp * rec(n+1, 1, np) / 2.0
	// 必勝法を使わずに引き分け
	ret += (1 - fp) * rec(n+1, 0, np) / 3.0

	saved[n][tgl][p] = true
	memo[n][tgl][p] = ret
	return ret
}

var p0, q int
var N = 100

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	p0, q = getI(), getI()
	ret := rec(0, 1, p0)/3.0 + 1/3.0
	out(ret)
}
