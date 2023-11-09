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
	L, R := getI(), getI()

	const size = int(1e6 + 1)
	p := make([]bool, size)
	q := make([]bool, R-L+1)

	for i := 2; i < size; i++ {
		p[i] = true
	}
	for i := 2; i < R-L+1; i++ {
		q[i] = true
	}

	for i := 2; i*i < size; i++ {
		for j := 2 * i; j < size; j += i {
			p[j] = false
		}
	}

	if R <= size {
		cnt := 0
		for i := L; i <= R; i++ {
			if p[i] {
				cnt++
			}
		}
		out(cnt)
		return
	}

	for i := 2; i*i <= R; i++ {
		if p[i] {
			pos := i * ((L + i - 1) / i)
			// out(i, "----", pos)
			for j := pos; j <= R; j += i {
				// out(j - L)
				q[j-L] = false
			}
		}
	}

	cnt := 0
	for i := 0; i < len(q); i++ {
		if q[i] {
			// out(i + L)
			cnt++
		}
	}
	out(cnt)
}
