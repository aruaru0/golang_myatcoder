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

const size = 1000000

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	P, Q := getI(), getI()

	a := make([]int, size+1)
	s := make([]int, size+1)
	m := make([]int, size+1)
	m[0] = P + 1
	for i := 1; i <= size; i++ {
		a[i] = P % i
		s[i] = a[i] + s[i-1]
		m[i] = P/(i+1) + 1
	}

	for i := 0; i < Q; i++ {
		l, r := getI(), getI()
		ans := 0
		if r <= size {
			ans = s[r] - s[l-1]
		} else {
			if l <= size {
				ans += s[size] - s[l-1]
				l = size + 1
			}
			x := P / l
			y := P / r
			for x != y {
				ne := m[x-1] - 1
				ans += P*(ne-l+1) - x*(ne+l)*(ne-l+1)/2
				x--
				l = m[x]
			}
			ans += P*(r-l+1) - y*(r+l)*(r-l+1)/2
		}
		out(ans)
	}
}
