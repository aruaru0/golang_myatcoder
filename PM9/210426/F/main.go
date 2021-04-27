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

var memo [210000][7]*bool

// true = takahashi win false = aoki win
func rec(i, v int) bool {
	v %= 7
	if i == n {
		return v == 0
	}
	if memo[i][v] != nil {
		return *memo[i][v]
	}
	var ret bool
	if x[i] == 'T' {
		if rec(i+1, v*10+int(s[i]-'0')) || rec(i+1, v*10) {
			// それ以外は勝つ方を選べる
			ret = true
		} else {
			// 以降の結果がどちらもfalseの場合は、負け確定（高橋）
			ret = false
		}
	} else {
		if rec(i+1, v*10+int(s[i]-'0')) && rec(i+1, v*10) {
			// 以降の結果、どうやってもtrueの場合は、負け確定（青木）
			ret = true
		} else {
			// それ以外は勝つ方を選べる
			ret = false
		}
	}
	memo[i][v] = &ret
	return ret
}

var n int
var s, x string

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n = getI()
	s = getS()
	x = getS()

	if rec(0, 0) {
		out("Takahashi")
		return
	}
	out("Aoki")
}
