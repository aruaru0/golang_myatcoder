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
	L, N1, N2 := getI(), getI(), getI()
	v1 := make([]int, N1)
	l1 := make([]int, N1)
	p1 := make([]int, N1)

	pos := 0
	for i := 0; i < N1; i++ {
		v, l := getI(), getI()
		v1[i], l1[i] = v, l
		p1[i] = pos
		pos += l
	}
	p1 = append(p1, L)

	v2 := make([]int, N2)
	l2 := make([]int, N2)
	p2 := make([]int, N2)
	pos = 0
	for i := 0; i < N2; i++ {
		v, l := getI(), getI()
		v2[i], l2[i] = v, l
		p2[i] = pos
		pos += l
	}
	p2 = append(p2, L)

	pos = 0
	cnt := 0
	for pos < L {
		idx1 := upperBound(p1, pos) - 1
		idx2 := upperBound(p2, pos) - 1
		p := min(p1[idx1]+l1[idx1], p2[idx2]+l2[idx2])
		// out("idx = ", idx1, idx2, "pos = ", pos, "v = ", v1[idx1], v2[idx2], "p=", p)
		if v1[idx1] == v2[idx2] {
			cnt += p - pos
			// out(p, pos, p-pos, "cnt = ", cnt)
		}
		pos = p
	}

	// out(v1, l1, p1)
	// out(v2, l2, p2)

	out(cnt)
}
