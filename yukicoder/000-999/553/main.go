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
	N := getF()
	r := make([]float64, int(N))
	for i := 0; i < int(N); i++ {
		r[i] = getF()
	}

	pre := 0.0
	x := 0.0
	y := 0.0
	for i := 1; i <= int(N); i++ {
		pre += math.Pow(2, r[i-1]/800) * math.Pow(0.9, float64(i))
		x += math.Pow(0.81, float64(i))
		y += math.Pow(0.9, float64(i))
	}
	Fn := math.Sqrt(x) / y
	fn := 1200 * (Fn - 0.229416) / (1.0 - 0.229416)
	// out(Fn, fn)

	if N == 1 {
		fn = 1200
	}
	X := 800 * math.Log2(pre/y)
	// out(X)
	ans := math.Max(0, X-fn)
	out(math.Round(ans))
}
