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
	N := getI()
	s := getS()

	a := make([]int, 0)
	for i := 0; i < N; i++ {
		if s[i] == '1' {
			a = append(a, i)
		}
	}

	k := len(a)
	// a の prefix sum を作る
	// prefix[i] = a[0] + a[1] + ... + a[i-1] (i番目は含まない)
	prefix := make([]int, k+1)
	for i := 0; i < k; i++ {
		prefix[i+1] = prefix[i] + a[i]
	}

	// sum_{j=0}^{k-1} |a[i] - a[j]| を O(1) で求める関数
	getSumAbsA := func(i int) int {
		// i より左側: (a[i] - a[j]) の和
		// i より右側: (a[j] - a[i]) の和
		leftCount := i
		rightCount := (k - 1) - i

		sumLeft := a[i]*leftCount - prefix[i]
		sumRight := (prefix[k] - prefix[i+1]) - a[i]*rightCount
		return sumLeft + sumRight
	}

	// sum_{j=0}^{k-1} |i - j| を O(1) で求める
	getSumAbsI := func(i int) int {
		// j < i の部分が i*(i+1)/2
		// j > i の部分が (k-1-i)*(k-i)/2
		left := i * (i + 1) / 2
		right := (k - 1 - i) * (k - i) / 2
		return left + right
	}

	ans := math.MaxInt
	for i := 0; i < k; i++ {
		valA := getSumAbsA(i)
		valI := getSumAbsI(i)
		tot := valA - valI
		if tot < ans {
			ans = tot
		}
	}
	out(ans)
}
