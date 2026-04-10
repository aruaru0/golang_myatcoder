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

func idx2rc(idx, N int) (int, int) {
	idx--
	return idx / N, idx % N
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	s, t := getI(), getI()
	p := getInts(M)
	p = append([]int{s}, p...)

	const inf = int(1e18)
	// s, t, m
	bit := 1 << (M + 1)
	// dp[S][i] : Sの頂点集合を通ってiに到着した時の最小コスト
	dp := make([][]int, bit)
	for i := 0; i < bit; i++ {
		dp[i] = make([]int, M+1)
		for j := 0; j < M+1; j++ {
			dp[i][j] = inf
		}
	}
	for i := 0; i < M+1; i++ {
		r0, c0 := idx2rc(p[0], N)
		r1, c1 := idx2rc(p[i], N)
		d := abs(r0-r1) + abs(c0-c1)
		dp[1|1<<i][i] = d
	}

	for b := 0; b < bit; b++ {
		for from := 0; from < M+1; from++ {
			if (b>>from)%2 == 0 { // fromが集合に含まれている
				continue
			}
			if dp[b][from] == inf {
				continue
			}
			for to := 0; to < M+1; to++ {
				if (b>>to)%2 == 1 { //　toが集合に含まれていない
					continue
				}
				r0, c0 := idx2rc(p[from], N)
				r1, c1 := idx2rc(p[to], N)
				d := abs(r0-r1) + abs(c0-c1)
				dp[b|(1<<to)][to] = min(dp[b|(1<<to)][to], dp[b][from]+d)
			}
		}
	}

	ans := inf
	for i := 0; i < M+1; i++ {
		r0, c0 := idx2rc(p[i], N)
		r1, c1 := idx2rc(t, N)
		d := abs(r0-r1) + abs(c0-c1)
		ans = min(ans, dp[bit-1][i]+d)
	}
	// out(dp[bit-1])
	out(ans)

}
