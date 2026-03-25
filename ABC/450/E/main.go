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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	x, y := getS(), getS()
	x = y + x

	S := []string{y, x}
	ll := []int{len(y), len(x)}
	cnt := make([][26]int, 2)
	for i := 0; i < 2; i++ {
		for j := 0; j < len(S[i]); j++ {
			cnt[i][int(S[i][j]-'a')]++
		}
	}
	csum := make([][]int, 26)
	for i := 0; i < 26; i++ {
		csum[i] = make([]int, len(x)+1)
	}
	for i := 0; i < len(x); i++ {
		csum[int(x[i]-'a')][i+1]++
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < len(x); j++ {
			csum[i][j+1] += csum[i][j]
		}
	}

	for ll[len(ll)-1] < int(1e18) {
		i := len(ll) - 1
		j := i - 1
		ll = append(ll, ll[i]+ll[j])
		cnt = append(cnt, cnt[i])
		n := len(cnt) - 1
		for c := 0; c < 26; c++ {
			cnt[n][c] += cnt[j][c]
		}
	}

	var f func(i, r, c int) int
	f = func(i, r, c int) int {
		if i <= 1 {
			return csum[c][r]
		}
		if ll[i-1] < r {
			return cnt[i-1][c] + f(i-2, r-ll[i-1], c)
		} else {
			return f(i-1, r, c)
		}
	}

	q := getI()
	for qi := 0; qi < q; qi++ {
		l, r, c := getI()-1, getI(), int(getS()[0]-'a')
		ans := f(len(ll), r, c) - f(len(ll), l, c)
		out(ans)
	}
}
