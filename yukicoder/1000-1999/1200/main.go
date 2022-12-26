package main

import (
	"bufio"
	"fmt"
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

func divisor(n int) []int {
	ret := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			ret = append(ret, n/i)
		}
	}
	return ret
}

func f(x, y int) {
	ret := 0
	if x < y {
		x, y = y, x
	}
	d := divisor(x - y)
	for _, i := range d {
		a := i + 1
		v, w := (x-y)/(a-1), (x+y)/(a+1)
		b := (v + w) / 2
		c := (w - v) / 2
		if b > 0 && c > 0 && (a-1)*(b-c) == (x-y) && (a+1)*(b+c) == (x+y) {
			ret++
		}
	}
	out(ret)
}

// 例外処理
func f2(x, y int) {
	ret := x - 1
	for i := 3; i*i <= x; i++ {
		if x%i == 0 {
			if i != x/i {
				ret += 2
			} else {
				ret++
			}
		}
	}
	if x%2 == 0 {
		if x/2 != 2 && x != 2 {
			ret++
		}
	}
	if x != 1 && x != 2 {
		ret++
	}
	out(ret)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getI()
	for i := 0; i < s; i++ {
		x, y := getI(), getI()
		if x == y {
			f2(x, y)
		} else {
			f(x, y)
		}
	}
}
