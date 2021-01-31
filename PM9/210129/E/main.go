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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()

	// 小さい時はコーナーケースを含めて全部試す
	// 2がコーナーケースっぽい
	if n < 100000 {
		cnt := 0
		for i := 2; i <= n; i++ {
			x := n
			for x%i == 0 {
				x /= i
			}
			if x%i == 1 {
				cnt++
			}
		}
		out(cnt)
		return
	}

	cnt := 2
	// n-1の約数なら１残せる
	for i := 2; i*i <= n; i++ {
		x := n - 1
		if x%i == 0 {
			cnt++
			if x/i != i {
				cnt++
			}
		}
	}
	// nの約数の場合は、割って１あまりが作れるか調べる
	for i := 2; i*i <= n; i++ {
		x := n
		if x%i == 0 {
			for x%i == 0 {
				x /= i
			}
			if x%i == 1 {
				cnt++
			}
			if n/i != i {
				y := n / i
				if n%y == 1 {
					cnt++
				}
			}
		}
	}
	out(cnt)
}
