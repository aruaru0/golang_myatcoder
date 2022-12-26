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

func pfs(n int, p map[int]int) map[int]int {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			p[i]++
			n /= i
		}
	}
	if n > 1 {
		p[n]++
	}
	return p
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	p := make(map[int]int)
	for i := 1; i <= N; i++ {
		p = pfs(i, p)
	}

	a := make([]int, 0)
	for _, e := range p {
		a = append(a, e)
	}
	sort.Ints(a)
	cnt := 0
	// 2 - 4 - 4
	x2 := len(a) - lowerBound(a, 2) - 2
	x4 := len(a) - lowerBound(a, 4)
	cnt += x2 * x4 * (x4 - 1) / 2

	// 2 24
	x2 = len(a) - lowerBound(a, 2) - 1
	x24 := len(a) - lowerBound(a, 24)
	cnt += x2 * x24
	// 4 14
	x4 = len(a) - lowerBound(a, 4) - 1
	x14 := len(a) - lowerBound(a, 14)
	cnt += x4 * x14
	// 74
	x74 := len(a) - lowerBound(a, 74)
	cnt += x74
	out(cnt)
}
