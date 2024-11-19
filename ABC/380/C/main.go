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
	N, K := getI(), getI()
	s := getS()

	one := make([][]byte, 0)
	zero := make([][]byte, 0)

	cur := byte('.')
	tmp := []byte{}
	for i := 0; i < N; i++ {
		if cur != s[i] {
			if cur == '1' {
				one = append(one, tmp)
			} else if cur == '0' {
				zero = append(zero, tmp)
			}
			tmp = []byte{}
		}
		cur = s[i]
		tmp = append(tmp, byte(s[i]))
	}

	if cur == '1' {
		one = append(one, tmp)
	} else {
		zero = append(zero, tmp)
	}

	idx0, idx1 := 0, 0
	t := []byte{}
	if s[0] == '0' {
		t = append(t, zero[idx0]...)
		idx0++
	}

	for i := 0; i < K-1; i++ {
		if idx1 != len(one) {
			t = append(t, one[idx1]...)
			idx1++
		}
		if i == K-2 {
			break
		}
		if idx0 != len(zero) {
			t = append(t, zero[idx0]...)
			idx0++
		}
	}

	t = append(t, one[idx1]...)
	idx1++
	t = append(t, zero[idx0]...)
	idx0++

	for len(one) > idx1 || len(zero) > idx0 {
		if idx0 != len(zero) {
			t = append(t, zero[idx0]...)
			idx0++
		}
		if idx1 != len(one) {
			t = append(t, one[idx1]...)
			idx1++
		}
	}

	// out(idx0, idx1)
	out(string(t))
}
