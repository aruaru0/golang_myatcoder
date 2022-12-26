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

// 先頭からの文字列と各文字から一致する文字列の長さを調べる
// 戻り値は各文字からの一致文字数（０は文字列の長さと一致）
func zalgo(str string) []int {
	n := len(str)
	a := make([]int, n)
	from, last := -1, -1
	a[0] = n
	for i := 1; i < n; i++ {
		idx := a[i]
		if from != -1 {
			idx = min(a[i-from], last-i)
			idx = max(0, idx)
		}
		for idx+i < n && str[idx] == str[idx+i] {
			idx++
		}
		a[i] = idx
		if last < i+idx {
			last = i + idx
			from = i
		}
	}
	return a
}

const mod = int(1e9 + 7)

var memo [11000]int

func rec(s string) int {
	if len(s) == 0 {
		return 1
	}
	if memo[len(s)] != 0 {
		return memo[len(s)]
	}

	// out(s, l, s[l:n-l])
	z := zalgo(s)
	res := 1

	for i := len(s) - 1; i >= 0; i-- {
		com := len(s) - i
		if z[i] == com && 2*com <= len(s) {
			res = (res + rec(s[com:len(s)-com])) % mod
		}
	}

	memo[len(s)] = res
	return res
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()
	ans := rec(s)
	out(ans)
}
