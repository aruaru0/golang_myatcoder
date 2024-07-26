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
	N := getI()
	if N <= 10 {
		out(N - 1)
		return
	}

	c := []int{0, 10, 9}
	x := 1
	for i := 0; x < 1e18; i++ {
		c = append(c, 9*x*10)
		if i%2 == 1 {
			x *= 10
		}
	}
	c = append(c, math.MaxInt64)

	// out(c, N)
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	pos := lowerBound(c, N)
	x = N - c[pos-1]
	x--
	// out(x, pos, pos/2)
	a := make([]int, 0)
	for i := 0; i < (pos+1)/2; i++ {
		a = append(a, x%10)
		x /= 10
	}
	a[len(a)-1]++
	// out(a)

	for i := 0; i < len(a); i++ {
		fmt.Fprint(wr, a[len(a)-1-i])
	}
	if pos%2 == 1 {
		a = a[1:]
	}
	for i := 0; i < len(a); i++ {
		fmt.Fprint(wr, a[i])
	}
	out()

	// 1001

	// 0
	// 1 2 3 4 5 6 7 8 9 9*1個
	// 11 22 33 44 55 66 77 88 99 ... 9個
	// 1x1 2x2 .... ->  90個
	// 1xx1, .. 9 * 9種 -> 90個
	// 1xxx1 2xxx2 3xxx3 ... 9*00~99-> 10個 -> 900個
	// 1xxxx1 2xxxx2 3xxxx3 ... 9*000~999 ->  900個
	// 9*10000個
	// 46個目は0を引いて45個
	// 9個引いて36個
	// 9個引いて27個
	// 90個中にいるので20なので3- 7なので6個目　結363

}
