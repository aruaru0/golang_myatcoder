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

var s []string

func calc0(ii, jj, k int) int {
	var x [4]int
	var y [4]int
	dx := []int{0, k, 0, k}
	dy := []int{0, 0, k, k}
	for i := 0; i < 4; i++ {
		x[i] = ii + dx[i]
		y[i] = jj + dy[i]
		if x[i] < 0 || x[i] >= 9 || y[i] < 0 || y[i] >= 9 {
			return 0
		}
	}
	ret := 1
	for i := 0; i < 4; i++ {
		if s[y[i]][x[i]] != '#' {
			ret = 0
		}
	}
	return ret
}

func calc(x, y, dx, dy int) int {
	var X [4]int
	var Y [4]int

	X[0] = x
	Y[0] = y

	X[1] = X[0] + dx
	Y[1] = Y[0] - dy

	X[2] = X[1] - dy
	Y[2] = Y[1] - dx

	X[3] = X[2] - dx
	Y[3] = Y[2] + dy

	for i := 0; i < 4; i++ {
		if X[i] < 0 || X[i] >= 9 || Y[i] < 0 || Y[i] >= 9 {
			return 0
		}
	}

	ret := 1
	for i := 0; i < 4; i++ {
		if s[Y[i]][X[i]] != '#' {
			ret = 0
		}
	}

	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s = make([]string, 9)
	for i := 0; i < 9; i++ {
		s[i] = getS()
	}
	ans := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				for l := 1; l < 9; l++ {
					if k == 0 && l == 0 {
						continue
					}
					if calc(i, j, k, l) != 0 {
						ans++
						// out(i, j, k, l)
					}
				}
			}
		}
	}
	out(ans)
}
