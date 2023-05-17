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

func f0(N int) int {
	return 1
}

func f1(N, M int) int {
	s := 0
	for i := 0; i < N; i++ {
		s++
	}
	for i := 0; i < M; i++ {
		s++
	}
	return s
}

func f2(N int) int {
	s := 0
	for i := 0; i < N; i++ {
		t := N
		cnt := 0
		for t > 0 {
			cnt++
			t /= 2
		}
		s += cnt
	}
	return s
}

func f3(N int) int {
	s := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s++
		}
	}
	return s
}

func f4(N int) int {
	s := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			s += i + j
		}
	}
	return s
}

func f5(N, M int) int32 {
	s := int32(0)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			s += int32(i + j)
		}
	}
	return s
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	N, M := getI(), getI()

	a0, a1, a2, a3, a4, a5 := -1, -1, -1, -1, -1, int32(-1)

	// 計算量が最も大きいもの1つだけコメントアウトする
	a0 = f0(N)
	a1 = f1(N, M)
	a2 = f2(N)
	a3 = f3(N)
	// a4 = f4(N);
	a5 = f5(N, M)

	out("f0:", a0)
	out("f1:", a1)
	out("f2:", a2)
	out("f3:", a3)
	out("f4:", a4)
	out("f5:", a5)
}
