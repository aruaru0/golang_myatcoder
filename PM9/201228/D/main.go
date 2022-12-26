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

type pair struct {
	a, b, i int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x := make([]pair, N)
	for i := 0; i < N; i++ {
		a, b := getI(), getI()
		x[i] = pair{a, b, i}
	}
	y := make([]pair, N)
	copy(y, x)
	sort.Slice(x, func(i, j int) bool {
		return x[i].a-x[j].b > x[j].a-x[i].b
	})
	// out(x)
	sort.Slice(y, func(i, j int) bool {
		return x[i].a-x[j].b < x[j].a-x[i].b
	})
	// out(y)

	eat := make([]bool, N)
	ta, ti := 0, 0
	ao, ai := 0, 0
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			for eat[x[ti].i] {
				ti++
			}
			// out("taka", x[ti])
			ta += x[ti].a
			eat[x[ti].i] = true
		} else {
			for eat[x[ai].i] {
				ai++
			}
			// out("aoki", x[ai])
			ao += x[ai].b
			eat[x[ai].i] = true
		}
	}
	out(ta - ao)
}
