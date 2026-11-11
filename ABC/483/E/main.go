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
	N, M := getI(), getI()
	s := make([][]byte, N)
	for i := 0; i < N; i++ {
		s[i] = []byte(getS())
	}

	toInt := func(s []byte) int {
		ret := 0
		for i := 0; i < len(s); i++ {
			ret = (ret << 8) + int(s[i])
		}
		return ret
	}

	m0 := make(map[int]int)
	for i := 0; i < N; i++ {
		m0[toInt(s[i])]++
	}

	m1 := make(map[int]int)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			backup := s[i][j]
			s[i][j] = '*'
			m1[toInt(s[i])]++
			s[i][j] = backup
		}
	}

	m2 := make(map[int]int)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			for k := j + 1; k < M; k++ {
				b0, b1 := s[i][j], s[i][k]
				s[i][j] = '*'
				s[i][k] = '*'
				m2[toInt(s[i])]++
				s[i][j] = b0
				s[i][k] = b1
			}

		}
	}

	x0, x1, x2 := 0, 0, 0
	for _, e := range m0 {
		x0 += e * (e - 1) / 2
	}
	for _, e := range m1 {
		x1 += e * (e - 1) / 2
	}
	for _, e := range m2 {
		x2 += e * (e - 1) / 2
	}

	out(x2 - (M-2)*x1 + (M*(M-1)/2-M+1)*x0)
}
