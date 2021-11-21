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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, T := getI(), getI()
	a := getInts(N)

	N1 := min(20, N)
	N2 := N - N1

	n := 1 << N1
	dp1 := make([]int, n)
	for bit := 0; bit < n; bit++ {
		for i := 0; i < N1; i++ {
			if (bit>>i)%2 == 0 {
				dp1[bit|(1<<i)] = dp1[bit] + a[i]
			}
		}
	}
	if N2 == 0 {
		ans := 0
		for i := 0; i < n; i++ {
			if dp1[i] <= T {
				ans = max(ans, dp1[i])
			}
		}
		out(ans)
		return
	}

	n2 := 1 << N2
	dp2 := make([]int, n2)
	for bit := 0; bit < n2; bit++ {
		for i := 0; i < N2; i++ {
			if (bit>>i)%2 == 0 {
				dp2[bit|(1<<i)] = dp2[bit] + a[N1+i]
			}
		}
	}
	sort.Ints(dp1)
	dp2 = append(dp2, inf)
	sort.Ints(dp2)

	// out(dp1)
	// out(dp2)
	ans := 0
	for i := 0; i < n; i++ {
		v := dp1[i]
		t := T - v
		if t < 0 {
			continue
		}
		pos := upperBound(dp2, t)
		ans = max(ans, v+dp2[max(0, pos-1)])
		// out("x", v, t, pos)

	}
	// out(N1, N2, dp2)
	out(ans)
}
