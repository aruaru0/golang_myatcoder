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

func primes(n int) []string {
	p := make([]bool, n+1)
	p[0] = true
	p[1] = true
	for i := 0; i <= n; i++ {
		if p[i] == false {
			for j := i * 2; j <= n; j += i {
				p[j] = true
			}
		}
	}
	ret := []string{}
	for i := 0; i <= n; i++ {
		if p[i] == false {
			ret = append(ret, strconv.Itoa(i))
		}
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()

	p := primes(9999999)

	make_pattern := func(s string) [][]int {
		m := make(map[byte][]int)
		for i := 0; i < len(s); i++ {
			m[s[i]] = append(m[s[i]], i)
		}
		x := make([][]int, 0)
		for _, e := range m {
			x = append(x, e)
		}
		sort.Slice(x, func(i, j int) bool {
			return x[i][0] < x[j][0]
		})

		return x
	}

	ms := make_pattern(s)

	check := func(t string) bool {
		// 同じパターンになるかチェック
		mt := make_pattern(t)
		if len(ms) != len(mt) {
			return false
		}
		for i := 0; i < len(ms); i++ {
			if len(ms[i]) != len(mt[i]) {
				return false
			}
			for j := 0; j < len(ms[i]); j++ {
				if ms[i][j] != mt[i][j] {
					return false
				}
			}
		}

		return true
	}

	for _, t := range p {
		if len(t) != len(s) {
			continue
		}
		if check(t) {
			out(t)
			return
		}
	}
	out(-1)
}
