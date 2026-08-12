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

var done [201][201][201]bool
var dp [201][201][201]float64

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, l := getI(), getI()
	a := getInts(n)

	ave := 0.0
	for _, e := range a {
		ave += float64(e)
	}
	ave /= float64(n)

	var f func(int, int, int) float64
	f = func(l, i2, i1 int) float64 {
		if done[l][i2][i1] {
			return dp[l][i2][i1]
		}
		if l == 0 {
			return 0
		}
		if i2 == 0 && i1 == 0 {
			return 0
		}
		res := 0.0
		p := 1.0 / float64(i2*2+i1)
		q := 1.0 / float64(i2*2+i1-1)

		if i1 != 0 { // 1
			res += (p * float64(i1)) * (f(l, i2, i1-1) + 1)
		}
		if i2 != 0 && i1 != 0 { // 2,1
			flg := 1.0
			if l == 1 {
				flg = 0.0
			}
			res += (p * float64(i2) * 2) * (q * float64(i1)) *
				(f(l-1, i2-1, i1) + flg)
		}
		if i2 != 0 { // 2,2 =
			res += (p * float64(i2) * 2) * (q * 1) *
				(f(l, i2-1, i1) + 1)
		}
		if i2 >= 2 { // 2,2 !=
			res += (p * float64(i2) * 2) * (q * float64(i2*2-2)) *
				(f(l-1, i2-2, i1+2))
		}
		done[l][i2][i1] = true
		dp[l][i2][i1] = res
		return res
	}

	ans := f(l, n, 0) * ave
	out(ans)
}
