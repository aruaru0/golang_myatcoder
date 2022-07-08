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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		pre[i] = getI() - 1
	}
	in := make([]int, n)
	for i := 0; i < n; i++ {
		in[i] = getI() - 1
	}
	ni := make([]int, n)
	for i := 0; i < n; i++ {
		ni[in[i]] = i
	}

	l := make([]int, n)
	r := make([]int, n)

	var f func(lp, li, n int) int
	f = func(lp, li, n int) int {
		if n == 0 {
			return 0
		}
		root := pre[lp]          // pre側のノード番号
		i := ni[root]            // in側のrootのインデックス
		if i < li || i >= li+n { // インデックスが現在のin側の参照より小さい、または、li+nより大きい（領域外の場合）
			return -1
		}
		ls := i - li     // 左側位置
		rs := n - 1 - ls // 右側位置
		l[root] = f(lp+1, li, ls)
		r[root] = f(lp+1+ls, li+ls+1, rs)
		if l[root] == -1 || r[root] == -1 {
			return -1
		}
		return root + 1
	}

	if f(0, 0, n) != 1 {
		out(-1)
		return
	}
	for i := 0; i < n; i++ {
		out(l[i], r[i])
	}
}
