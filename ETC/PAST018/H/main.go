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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
	}
	a := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		a[i] = make([]int, N+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			x := 0
			if s[i][j] == '#' {
				x = 1
			}
			a[i+1][j+1] += x
		}
	}

	for i := 0; i <= N; i++ {
		for j := 0; j < N; j++ {
			a[i][j+1] += a[i][j]
		}
	}
	for i := 0; i <= N; i++ {
		for j := 0; j < N; j++ {
			a[j+1][i] += a[j][i]
		}
	}

	// for i := 0; i <= N; i++ {
	// 	out(a[i])
	// }

	calc := func(x0, y0, x1, y1 int) int {
		an := a[x0][y0]
		bn := a[x0][y1]
		cn := a[x1][y0]
		dn := a[x1][y1]
		return dn - bn - cn + an
	}

	ans := 0
	for n := 3; n <= N; n++ {
		for i := 0; i <= N-n; i++ {
			for j := 0; j <= N-n; j++ {
				x0 := calc(i, j, i+n, j+1)
				x1 := calc(i, j, i+1, j+n)
				x2 := calc(i+n-1, j, i+n, j+n)
				y0 := calc(i+1, j+1, i+n-1, j+n)
				if x0 == n && x1 == n && x2 == n && y0 == 0 {
					// out(i, j, "n=", n)
					ans = max(ans, n-2)
				}
			}
		}
	}
	out(ans)
}
