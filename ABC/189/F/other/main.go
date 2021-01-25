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

type L struct {
	p, e float64
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K := getI(), getI(), getI()
	a := make(map[int]bool)
	for i := 0; i < K; i++ {
		a[getI()] = true
	}
	m := float64(M)
	dp := make([]L, N+M+1)
	dpsum := make([]L, N+M+1)
	holesum := L{0, 0}
	anssum := L{0, 0}
	dp[0] = L{1, 0}
	dpsum[0] = dp[0]

	for i := 1; i <= N+M; i++ {
		tmp := dpsum[i-1]
		if i-M-1 >= 0 {
			tmp.p -= dpsum[i-M-1].p
			tmp.e -= dpsum[i-M-1].e
		}
		tmp.p /= m
		tmp.e /= m
		tmp.e += tmp.p
		if a[i] {
			holesum.p += tmp.p
			holesum.e += tmp.e
			tmp = L{0, 0}
		}
		if i >= N {
			anssum.p += tmp.p
			anssum.e += tmp.e
			tmp = L{0, 0}
		}
		dp[i] = tmp
		dpsum[i].p = dpsum[i-1].p + dp[i].p
		dpsum[i].e = dpsum[i-1].e + dp[i].e
		// out(dp, anssum, holesum)
	}

	if anssum.p == 0 {
		out(-1)
		return
	}
	ans := anssum.e / anssum.p
	if holesum.p != 0 {
		ans += holesum.e / (1 - holesum.p)
	}
	out(ans)
}
