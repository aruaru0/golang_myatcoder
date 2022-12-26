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

func f(n, x, k int, a []int) int {
	b := make([]int, n)
	for j := 0; j < n; j++ {
		b[j] = abs(a[j] - x)
	}
	sort.Ints(b)
	ans := 0
	for i := 0; i < n; i++ {
		if i < k {
			ans += b[i]
		} else {
			ans -= b[i]
		}
	}
	return ans
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	A := getInts(N)
	sum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		sum[i] += sum[i-1] + A[i-1]
	}
	ans := inf
	l, r := 0, K
	// 処理は理解できるが実装できる自身がないかも。
	for i := 0; i < N; i++ {
		if i == r {
			l++
			r++
		}
		for r < N && abs(A[i]-A[l]) > abs(A[i]-A[r]) {
			l++
			r++
		}
		now := A[i]*(i-l) - (sum[i] - sum[l]) + (sum[r] - sum[i]) - A[i]*(r-i)
		now -= (sum[N] - sum[r]) - A[i]*(N-r) + A[i]*l - sum[l]
		ans = min(ans, now)
	}
	out(ans)
}
