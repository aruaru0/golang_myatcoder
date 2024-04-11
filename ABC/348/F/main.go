package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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

func getI16() int16 {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return int16(i)
}

func getInts16(N int) []int16 {
	ret := make([]int16, N)
	for i := 0; i < N; i++ {
		ret[i] = getI16()
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = getInts(m)
	}
	b := make([][]*big.Int, m)
	for j := 0; j < m; j++ {
		b[j] = make([]*big.Int, 1000)
		for v := 0; v < 1000; v++ {
			b[j][v] = big.NewInt(0)
		}
		for i := 0; i < n; i++ {
			v := a[i][j]
			b[j][v] = b[j][v].SetBit(b[j][v], i, 1)
		}
	}

	// i番目の数列と一致する数
	dp := make([]*big.Int, n)
	for i := range dp {
		dp[i] = big.NewInt(0)
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			dp[i] = dp[i].Xor(dp[i], b[j][a[i][j]])
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if dp[i].Bit(j) == 1 {
				ans++
			}
		}
	}
	out(ans)
}
