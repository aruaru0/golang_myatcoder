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

type box struct {
	w, v int
}

var M, N, Q int
var b []box
var x []int

func solve() {
	l, r := getI()-1, getI()-1
	a := make([]int, 0)
	for i := 0; i < M; i++ {
		if l <= i && i <= r {
			continue
		}
		a = append(a, x[i])
	}
	sort.Ints(a)

	used := make([]bool, len(a))
	ans := 0
	for i := 0; i < N; i++ {
		pos := lowerBound(a, b[i].w)
		for pos < len(a) && used[pos] == true {
			pos++
		}
		if pos == len(a) {
			continue
		}
		used[pos] = true
		ans += b[i].v
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, Q = getI(), getI(), getI()
	b = make([]box, N)
	for i := 0; i < N; i++ {
		b[i] = box{getI(), getI()}
	}
	sort.Slice(b, func(i, j int) bool {
		if b[i].v == b[j].v {
			return b[i].w < b[j].w
		}
		return b[i].v > b[j].v
	})
	x = getInts(M)
	for i := 0; i < Q; i++ {
		solve()
	}
}
