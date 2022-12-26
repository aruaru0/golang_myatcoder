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

type pair struct {
	pos, diff, cost int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, C := getI(), getI()

	xs := make([]int, N+2)
	vs := make([]int, N+2)
	sumv := make([]int, N+2)
	for i := 1; i < N+1; i++ {
		xs[i], vs[i] = getI(), getI()
		sumv[i] = sumv[i-1] + vs[i]
	}
	xs[N+1] = C
	vs[N+1] = 0
	sumv[N+1] = sumv[N]

	lmax := make([]int, N+2)
	lmax2 := make([]int, N+2)
	for left := 1; left < N+1; left++ {
		v := sumv[left]
		d := xs[left]
		lmax[left] = max(lmax[left-1], v-d)
		lmax2[left] = max(lmax2[left-1], v-2*d)
	}
	rmax := make([]int, N+2)
	rmax2 := make([]int, N+2)
	for right := N; right > 0; right-- {
		v := sumv[N] - sumv[right-1]
		d := C - xs[right]
		rmax[right] = max(rmax[right-1], v-d)
		rmax2[right] = max(rmax2[right-1], v-2*d)
	}

	ans := 0
	for i := 0; i <= N; i++ {
		c1 := lmax2[i] + rmax[i+1]
		c2 := lmax[i] + rmax2[i+1]
		ans = nmax(ans, c1, c2)
	}
	out(ans)
}
