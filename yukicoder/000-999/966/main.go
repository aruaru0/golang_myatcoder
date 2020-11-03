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

func f(a, b, c int) int {
	ret := 0
	// out("-----")
	for {
		// out(a, b, c)
		if a <= 0 || b <= 0 || c <= 0 {
			return -1
		}
		if a > b && c > b && a != c {
			return ret
		}
		if a < b && c < b && a != c {
			return ret
		}

		switch {
		case a == b && b == c:
			a -= 2
			c--
			ret += 3
		case a == c:
			ret++
			a--
		case a == b && b > c:
			a--
			ret++
		case a == b && b < c:
			b--
			ret++
		case c == b && b > a:
			c--
			ret++
		case c == b && b < a:
			b--
			ret++
		case a > b:
			x := a - b + 1
			y := b - c + 1
			if x >= y && b-y != 0 {
				ret += y
				b -= y
			} else {
				ret += x
				a -= x
			}
		case c > b:
			x := c - b + 1
			y := b - a + 1
			if x >= y && b-y != 0 {
				ret += y
				b -= y
			} else {
				ret += x
				c -= x
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		a, b, c := getI(), getI(), getI()
		out(f(a, b, c))
	}
}
