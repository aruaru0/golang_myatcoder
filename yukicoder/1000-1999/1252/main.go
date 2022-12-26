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

func solve() {
	d, A, B := getI(), getI(), getI()
	n := A
	tot := 0
	for n > 0 {
		tot += n % d
		tot = tot/d + tot%d
		n /= d
	}
	start := tot
	n = B
	tot = 0
	for n > 0 {
		tot += n % d
		tot = tot/d + tot%d
		n /= d
	}
	end := tot
	diff := (B - A + 1)
	if start != 1 {
		diff -= d - start
	}
	if end != d-1 {
		diff -= end
	}
	x := diff / (d - 1)
	// out(x, "start", start, "end", end)
	// out((d-start)*(start+d-1)/2,
	// 	end*(1+end)/2)
	ans := x * (d - 1) * (1 + d - 1) / 2
	// out("ans", ans)
	if start != 1 {
		ans += (d - start) * (start + d - 1) / 2
	}
	if end != d-1 {
		ans += end * (1 + end) / 2
	}
	out(ans)
	// out(ans, "-----------------")
	// out("xxxx")
	// ans = 0
	// for i := A; i <= B; i++ {
	// 	n := i
	// 	tot := 0
	// 	for n > 0 {
	// 		tot += n % d
	// 		tot = tot/d + tot%d
	// 		n /= d
	// 	}
	// 	ans += tot
	// 	out(tot, ans)
	// }
	// out(ans, "-----------------")
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		solve()
	}
}
