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

func check(r1, c1, r2, c2 int) int {
	if r1 == r2 && c1 == c2 {
		return 0
	}

	if abs(r1-r2) == abs(c1-c2) {
		return 1
	}

	if (r1+c1)%2 == (r2+c2)%2 {
		return 2
	}

	return 3
}

func check2(r1, c1, r2, c2 int) int {
	for y := -2; y <= 2; y++ {
		for x := -2; x <= 2; x++ {
			if r1+y == r2 && c1+x == c2 {
				return 1
			}
		}
	}
	if r1+3 == r2 && c1 == c2 {
		return 1
	}
	if r1-3 == r2 && c1 == c2 {
		return 1
	}
	if r1 == r2 && c1+3 == c2 {
		return 1
	}
	if r1 == r2 && c1-3 == c2 {
		return 1
	}
	return 3
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	r1, c1 := getI(), getI()
	r2, c2 := getI(), getI()
	if r1 == r2 && c1 == c2 {
		out(0)
		return
	}
	ans := 3
	for y := -2; y <= 2; y++ {
		for x := -2; x <= 2; x++ {
			ret := check(r1+y, c1+x, r2, c2)
			ret2 := check2(r1+y, c1+x, r2, c2)
			ret = min(ret, ret2)
			// out(x, y, r1+y, c1+x, r2, c2, ret)
			if y != 0 || x != 0 {
				ret++
			}
			ans = min(ans, ret)
		}
	}
	ans = min(ans, check(r1+3, c1, r2, c2)+1)
	ans = min(ans, check(r1-3, c1, r2, c2)+1)
	ans = min(ans, check(r1, c1+3, r2, c2)+1)
	ans = min(ans, check(r1, c1-3, r2, c2)+1)

	out(ans)
}
