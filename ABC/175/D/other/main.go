package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	// fmt.Println(x...)
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

const inf = math.MaxUint64 >> 1

var N, K int
var p []int
var c []int

func solveShort(s int) int {
	n := p[s]
	sum := 0
	ret := -inf
	for i := 0; i < K; i++ {
		sum += c[n]
		n = p[n]
		ret = max(ret, sum)
	}
	return ret
}

func solve(s int) int {
	// search loop
	n := p[s]
	used := make([]bool, N)
	r := make([]int, 0)
	for used[n] == false {
		r = append(r, n)
		used[n] = true
		n = p[n]
	}
	// pre
	if K < len(r) {
		m := -inf
		sum := 0
		for i := 0; i < K; i++ {
			sum += c[r[i]]
			m = max(m, sum)
		}
		return m
	}
	ret := -inf
	f := 0
	sum := 0
	for r[f] != n {
		sum += c[r[f]]
		ret = max(ret, sum)
		f++
	}
	// loop
	loop := 0
	mm := ret
	for i := f; i < len(r); i++ {
		loop += c[r[i]]
		mm = max(mm, sum+loop)
	}
	if loop < 0 {
		return mm
	}
	loopsize := len(r) - f
	nn := (K - f) / loopsize
	re := (K - f) % loopsize
	if nn > 0 {
		re += loopsize
		nn--
	}
	out(nn, re)
	sum += loop * nn
	ret = sum
	out(f, r)
	cur := r[f]
	out(sum)
	for i := 0; i < re; i++ {
		sum += c[cur]
		out(sum, cur)
		cur = p[cur]
		ret = max(ret, sum)
	}
	out(ret)
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K = getInt(), getInt()
	p = make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getInt() - 1
	}
	c = getInts(N)

	ans := -inf
	for i := 0; i < N; i++ {
		var ret int
		// if K < 10000 {
		// 	ret = solveShort(i)
		// } else {
		ret = solve(i)
		// }
		ans = max(ans, ret)
	}
	fmt.Println(ans)
}
