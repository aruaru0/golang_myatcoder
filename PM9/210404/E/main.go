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

func f(x []int) int {
	ret := 25
	cnt := 25
	for i := 0; i < len(x); i++ {
		// out(x[i], cnt, ret)
		if x[i] == 0 {
			cnt++
			continue
		}
		ret = min(ret, cnt)
		cnt = 1
	}
	// out(ret)
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	d := make([]int, 13)
	d[0] = 1
	for i := 0; i < N; i++ {
		t := getI()
		d[t]++
		if d[t] > 2 {
			out(0)
			return
		}
	}

	if d[0] >= 2 || d[12] >= 2 {
		out(0)
		return
	}

	n := 1 << 11
	ans := 0
	for i := 0; i < n; i++ {
		x := make([]int, 25)
		x[0] = d[0]
		x[24] = d[0]
		x[12] = d[12]
		for j := 0; j <= 11; j++ {
			if d[j] == 2 {
				x[j] = 1
				x[24-j] = 1
				continue
			}
			if (i>>j)%2 == 0 {
				x[j] = d[j]
			} else {
				x[24-j] = d[j]
			}
		}
		// out(x)
		ans = max(ans, f(x))
	}
	out(ans)
}
