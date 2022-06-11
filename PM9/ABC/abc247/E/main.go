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

func f(a []int, X, Y int) int {
	ret := 0
	n := len(a)
	l, r := 0, 0
	cntX, cntY := 0, 0
	for l < n {
		for r < n && (cntX == 0 || cntY == 0) {
			if a[r] == X {
				cntX++
			}
			if a[r] == Y {
				cntY++
			}
			r++
		}
		if cntX > 0 && cntY > 0 {
			ret += n - r + 1
		}
		if a[l] == X {
			cntX--
		}
		if a[l] == Y {
			cntY--
		}
		l++
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, X, Y := getI(), getI(), getI()
	a := getInts(N)
	a = append(a, 1e18)
	ans := 0
	for i := 0; i < N; {
		b := make([]int, 0)
		for {
			if a[i] < Y || a[i] > X {
				i++
				break
			}
			b = append(b, a[i])
			i++
		}
		if len(b) == 0 {
			continue
		}
		ans += f(b, X, Y)
	}
	out(ans)
}
