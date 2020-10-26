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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	t := make([]int, N)
	for i := 0; i < N; i++ {
		t[i] = -1
	}
	b := make([]int, 0)
	s := make([]int, 0)
	s = append(s, 0)
	x := 0
	start := 0
	for i := 0; ; i++ {
		r := x % N
		if t[r] != -1 {
			start = t[r]
			break
		}
		t[r] = i
		b = append(b, x)
		s = append(s, a[r])
		x += a[r]
	}
	for i := 1; i < len(s); i++ {
		s[i] += s[i-1]
	}

	// out(start, b)
	// out(s)
	Q := getI()
	for i := 0; i < Q; i++ {
		k := getI()
		if len(b) > k {
			out(s[k])
		} else {
			k -= len(b)
			// out(s[len(s)-1], s[start])
			sum := s[len(s)-1] - s[start]
			n := k / (len(s) - 1 - start)
			m := k % (len(s) - 1 - start)
			rest := s[start+m] - s[start]
			// out(sum, k, n, m)
			out(s[len(s)-1] + n*sum + rest)
		}
	}
}

// [3 7 9 16 19 26 29 36 39 46 49 56 59 66 69 76 79 86 89 96]
