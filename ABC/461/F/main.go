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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type Pair struct {
	c, s int
}

const mod = 998244353

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	ds := make([]int, 0)
	{
		s := make(map[int]bool)
		for i := 1; i*i <= N; i++ {
			if N%i == 0 {
				s[i] = true
				s[N/i] = true
			}
		}
		for e := range s {
			ds = append(ds, e)
		}
		sort.Ints(ds)
	}

	m := 13
	dp := make([]map[int]Pair, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make(map[int]Pair)
	}
	dp[0][1] = Pair{1, 0}
	dp[1][1] = Pair{1, 1}
	for _, d := range ds {
		if d > 1 {
			for i := m - 1; i >= 0; i-- {
				for p, cs := range dp[i] {
					c, s := cs.c, cs.s
					if (N/p)%d != 0 {
						continue
					}
					np := p * d
					nextPair := dp[i+1][np]
					nextPair.c = (nextPair.c + c) % mod
					nextPair.s = (nextPair.s + s + (c*d)%mod) % mod
					dp[i+1][np] = nextPair
				}
			}
		}
	}

	fac := make([]int, m+1)
	for i := 0; i < m+1; i++ {
		fac[i] = 1
	}
	for i := 0; i < m; i++ {
		fac[i+1] = (fac[i] * (i + 1)) % mod
	}
	ans := 0
	for i := 1; i <= m; i++ {
		ans += (dp[i][N].s * fac[i]) % mod
		ans %= mod
	}
	out(ans)
}
