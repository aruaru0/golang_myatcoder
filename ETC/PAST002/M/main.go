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
	D, L, N := getI(), getI(), getI()

	C := getInts(D)
	C = append(C, C...) // 2周分

	Q := make([][3]int, N)
	for i := 0; i < N; i++ {
		Q[i] = [3]int{getI(), getI(), getI()}
	}

	const MaxC = 100001
	A := make([][]int, MaxC)
	B := make([][]int, MaxC)
	for d := 0; d < 2*D; d++ {
		c := C[d]
		A[c] = append(A[c], d)
		if len(B[c]) == 0 {
			B[c] = append(B[c], 0)
		} else {
			last := B[c][len(B[c])-1]
			interval := (d - A[c][len(A[c])-2] - 1) / L
			B[c] = append(B[c], interval+1+last)
		}
	}

	T := make([]int, MaxC)
	for c := 0; c < MaxC; c++ {
		if len(A[c]) > 0 {
			T[c] = B[c][len(A[c])/2]
		}
	}

	ans := func(k, f, t int) int {
		if len(A[k]) == 0 {
			return 0
		}
		f--
		i := lowerBound(A[k], f)
		t0 := (A[k][i] - f + L - 1) / L
		if t0 >= t {
			return 0
		}
		t -= t0
		if i >= len(A[k])/2 {
			i -= len(A[k]) / 2
		}
		q := t / T[k]
		res := q * (len(A[k]) / 2)
		t %= T[k]
		if t == 0 {
			return res
		}
		j := upperBound(B[k], B[k][i]+t-1)
		res += j - i
		return res
	}

	for _, q := range Q {
		out(ans(q[0], q[1], q[2]))
	}
}
