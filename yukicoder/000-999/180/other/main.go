package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

const inf = int(1e18)

var N int
var a, b []int

func solve(f, step int) (int, int) {

	diff := inf
	ans := 0
	pre := f
	for x := pre; ; x += step {
		mi, ma := inf, 0
		for i := 0; i < N; i++ {
			c := a[i] + b[i]*x
			mi = min(mi, c)
			ma = max(ma, c)
		}
		d := ma - mi
		// out(mi, ma, x, d)
		if d < diff {
			diff = d
			pre = ans
			ans = x
		} else {
			break
		}
	}
	return pre, ans
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	a = make([]int, N)
	b = make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getInt(), getInt()
	}

	a, b := 0, 0
	for n := int(1e9); n > 0; n /= 10 {
		a, b = solve(a, n)
		// out(a, b, n)
	}
	if b == 0 {
		out(1)
		return
	}
	out(b)
}
