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

var B, N int
var c []int
var ans int

func calc(p int) int {
	b := B
	cnt := 0
	for i := 0; i < N; i++ {
		cnt += abs(p - c[i])
		b += c[i] - p
	}
	if b < 0 {
		return inf
	}
	ans = min(ans, cnt)
	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)
	B, N = getInt(), getInt()
	c = getInts(N)

	ans = inf
	l := 0
	r := B
	for _, e := range c {
		r += e
	}
	r /= N
	for r-l > 10 {
		m0 := (l + l + r) / 3
		m1 := (l + r + r) / 3
		if calc(m0) <= calc(m1) {
			r = m1
		} else {
			l = m0
		}
	}

	for i := max(0, l-100); i < l+100; i++ {
		calc(i)
	}
	out(ans)
}
