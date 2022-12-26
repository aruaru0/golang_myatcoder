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

func slideMin(a []int, k int) []int {
	n := len(a)
	b := make([]int, n-k+1)
	s, t := 0, 0

	deq := make([]int, n)
	for i := 0; i < n; i++ {
		for s < t && a[deq[t-1]] >= a[i] {
			t--
		}
		deq[t] = i
		t++
		if i-k+1 >= 0 {
			b[i-k+1] = a[deq[s]]
			if deq[s] == i-k+1 {
				s++
			}
		}
	}
	return b
}

func slideMax(a []int, k int) []int {
	n := len(a)
	b := make([]int, n-k+1)
	s, t := 0, 0

	deq := make([]int, n)
	for i := 0; i < n; i++ {
		for s < t && a[deq[t-1]] <= a[i] {
			t--
		}
		deq[t] = i
		t++
		if i-k+1 >= 0 {
			b[i-k+1] = a[deq[s]]
			if deq[s] == i-k+1 {
				s++
			}
		}
	}
	return b
}

type pair struct {
	a, b int
}

func f(a []int, k int) {
	n := len(a)
	for i := 0; i+k <= n; i++ {
		b := make([]int, n)
		copy(b, a)
		sort.Ints(b[i : i+k])
		out(b)
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	p := getInts(N)

	df := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		if p[i+1] > p[i] {
			df[i] = 1
		} else {
			df[i] = 0
		}
	}
	tot := make([]int, N)
	for i := 0; i < N-1; i++ {
		tot[i+1] = tot[i] + df[i]
	}
	cnt := 0
	for i := K - 1; i < N; i++ {
		if tot[i]-tot[i-K+1] == K-1 {
			cnt++
			break
		}
	}
	minP := slideMin(p, K)
	maxP := slideMax(p, K)
	for i := 0; i < N-K+1; i++ {
		if tot[i+K-1]-tot[i] == K-1 {
			continue
		}
		if minP[i] == p[i] && i+K < N && maxP[i+1] == p[i+K] {
			continue
		}
		cnt++
	}
	out(cnt)
}
