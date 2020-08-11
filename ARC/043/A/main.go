package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

func getFloat() float64 {
	return float64(getInt())
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, A, B := getInt(), getFloat(), getFloat()
	s := make([]float64, N)
	sum := 0.0
	smax := float64(0.0)
	smin := float64(1e10)
	for i := 0; i < N; i++ {
		s[i] = getFloat()
		sum += s[i]
		smax = math.Max(smax, s[i])
		smin = math.Min(smin, s[i])
	}
	ave := sum / float64(N)
	if smax-smin == 0 {
		out(-1)
		return
	}
	P := B / (smax - smin)
	Q := A - P*ave
	// out(ave, smax, smin)
	out(P, Q)
}
