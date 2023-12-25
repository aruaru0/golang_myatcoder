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
	N, K := getI(), getI()
	a := getInts(K)

	n := make(map[int]int)
	for i := 0; i < N; i++ {
		n[i+1] = 2
	}
	for i := 0; i < K; i++ {
		n[a[i]]--
	}

	p := make([]int, 0)
	for e, v := range n {
		for i := 0; i < v; i++ {
			p = append(p, e)
		}
	}

	sort.Ints(p)

	front := make([]int, 0)
	for i := 1; i < len(p); i += 2 {
		front = append(front, p[i]-p[i-1])
	}
	back := make([]int, len(front))
	pos := len(back) - 1
	for i := len(p) - 1; i > 0; i -= 2 {
		back[pos] = p[i] - p[i-1]
		pos--
	}

	for i := 0; i < len(front)-1; i++ {
		front[i+1] += front[i]
	}
	for i := len(back) - 1; i > 0; i-- {
		back[i-1] += back[i]
	}

	ans := min(front[len(front)-1], back[0])
	for i := 0; i < len(front)-1; i++ {
		ans = min(ans, front[i]+back[i+1])
	}
	out(ans)
}
