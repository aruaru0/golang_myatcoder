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

func f(l, r int) int {
	if r > l {
		r--
	}
	// out("R", r)
	return r
}

func solve(n int) {
	ans := 0
	// 2
	if n > 100 {
		ans += 10 - 1
	} else {
		l := n % 10
		r := n / 10
		ans += f(l, r)
		out(ans)
		return
	}
	// 4
	if n > 10000 {
		ans += 100 - 10
	} else {
		l := n % 100
		r := n / 100
		ans += max(0, f(l, r)-10+1)
		out(ans)
		return
	}
	// 6
	if n > 1000000 {
		ans += 1000 - 100
	} else {
		l := n % 1000
		r := n / 1000
		ans += max(0, f(l, r)-100+1)
		out(ans)
		return
	}
	// 8
	if n > 100000000 {
		ans += 10000 - 1000
	} else {
		l := n % 10000
		r := n / 10000
		ans += max(0, f(l, r)-1000+1)
		out(ans)
		return
	}
	// 10
	if n > 10000000000 {
		ans += 100000 - 10000
	} else {
		l := n % 100000
		r := n / 100000
		ans += max(0, f(l, r)-10000+1)
		out(ans)
		return
	}
	// 12
	if n > 1000000000000 {
		ans += 1000000 - 100000
	} else {
		l := n % 1000000
		r := n / 1000000
		ans += max(0, f(l, r)-100000+1)
		out(ans)
		return
	}

	out(ans)
}

func solve2(n int) {
	cnt := 0
	for i := 0; i <= n; i++ {
		s := strconv.FormatInt(int64(i), 10)
		if len(s)%2 != 0 {
			continue
		}
		if s[:len(s)/2] == s[len(s)/2:] {
			cnt++
			out(s)
		}
	}
	out(cnt)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()

	solve(n)
}
