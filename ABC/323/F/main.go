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
	ax, ay, bx, by, cx, cy := getI(), getI(), getI(), getI(), getI(), getI()

	ax -= cx
	bx -= cx
	ay -= cy
	by -= cy
	if bx < 0 {
		ax *= -1
		bx *= -1
	}
	if by < 0 {
		ay *= -1
		by *= -1
	}
	if by == 0 {
		ax, ay = ay, ax
		bx, by = by, bx
	}

	dist := func(x, y int) int {
		if ax == x && ay == y {
			return 0
		}
		res := abs(ax-x) + abs(ay-y)
		if ax == x && ax == bx {
			if (ay < by) != (y < by) {
				res += 2
			}
		}
		if ay == y && ay == by {
			if (ax < bx) != (x < bx) {
				res += 2
			}
		}
		return res
	}

	if bx == 0 {
		ans := dist(bx, by+1) + by
		out(ans)
		return
	}

	ans := min(dist(bx, by+1), dist(bx+1, by))
	ans += bx + by + 2
	out(ans)
}
