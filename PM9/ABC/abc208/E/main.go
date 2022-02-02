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

// DPであることは理解できるが、遷移の書き方がよくわからない
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, k := getI(), getI()
	n++

	dp := make(map[int]int)
	// nを桁に分解して、上の桁からに並べなおす
	ds := make([]int, 0)
	for n != 0 {
		ds = append(ds, n%10)
		n /= 10
	}
	for i := 0; i < len(ds)/2; i++ {
		ds[i], ds[len(ds)-1-i] = ds[len(ds)-1-i], ds[i]
	}

	//
	x := 1
	first := true
	for _, d := range ds {
		pre := make(map[int]int)
		dp, pre = pre, dp
		// f = 積　s = 場合の数
		for f, s := range pre {
			// 以下の場合
			for i := 0; i < 10; i++ {
				dp[f*i] += s
			}
		}
		for i := 0; i < d; i++ {
			if i == 0 && first { // 先頭の桁に0は入らない
				continue
			}
			dp[x*i]++
		}
		if !first {
			for i := 1; i <= 9; i++ {
				dp[i]++
			}
		}
		x *= d
		first = false
	}

	ans := 0
	for f, s := range dp {
		if f <= k {
			ans += s
		}
	}
	out(ans)
}
