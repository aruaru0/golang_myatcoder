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
	s := getS()

	n := len(s)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = -1
	}
	t := make([]int, 0)
	// pに括弧の対応位置を入れる
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			t = append(t, i)
		} else if s[i] == ')' {
			x := t[len(t)-1]
			t = t[:len(t)-1]
			p[x] = i
			p[i] = x
		}
	}

	s += ")"
	var f func(i int, mode byte)
	f = func(i int, mode byte) {
		if mode == 'R' { // 現在が(が開いている場合
			if s[i] == ')' { // )がきたら終わり
				return
			} else if s[i] == '(' { // さらに開いたら
				f(p[i]-1, 'L') // )モードで１つ前探索し
				f(p[i]+1, 'R') // (モードで)
			} else {
				fmt.Fprint(wr, string(s[i]))
				f(i+1, 'R')
			}
		} else {
			if s[i] == '(' {
				return
			} else if s[i] == ')' {
				f(p[i]+1, 'R')
				f(p[i]-1, 'L')
			} else {
				v := s[i]
				v = v ^ 32
				fmt.Fprint(wr, string(v))
				f(i-1, 'L')
			}
		}
	}
	f(0, 'R')
	out()
}
