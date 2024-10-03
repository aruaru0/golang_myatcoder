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

func lowerBoundP(a []pair, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].v >= x
	})
	return idx
}

type pair struct {
	v, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K := getI(), getI(), getI()
	v := 0
	a := make([]pair, N)
	for i := 0; i < N; i++ {
		x := getI()
		a[i] = pair{x, i}
		v += x
	}
	v = K - v

	sort.Slice(a, func(i, j int) bool {
		if a[i].v == a[j].v {
			return a[i].idx < a[j].idx
		}
		return a[i].v < a[j].v
	})

	s := make([]int, N+1)
	for i := 0; i < N; i++ {
		s[i+1] = s[i] + a[i].v
	}
	ans := make([]int, N)

	// 全員当確の場合は例外処理
	if M == N {
		outSlice(ans)
		return
	}

	// l, rの範囲に入るのにどれだけ必要か
	need := func(l, r, x int) int {
		pos := lowerBoundP(a[l:r], x) + l
		// out(pos, a, l, r, x)
		return (pos-l)*x - (s[pos] - s[l])
	}

	for i := 0; i < N; i++ {
		ng, ok := -1, v+1
		for ng+1 != ok {
			m := (ok + ng) / 2
			x := a[i].v + m + 1
			vc := 0
			if i < N-M {
				vc = need(N-M, N, x)
			} else {
				vc = (i-(N-M-1))*x - (s[i] - s[N-M-1])
				vc += need(i+1, N, x)
			}
			if vc > v-m {
				ok = m
			} else {
				ng = m
			}
		}
		if ok > v {
			ans[a[i].idx] = -1
		} else {
			ans[a[i].idx] = ok
		}
	}

	outSlice(ans)
}
