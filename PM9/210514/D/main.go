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

func calcScore(s string, x int) int {
	a := make(map[int]int)
	for i := 0; i < 4; i++ {
		v := int(s[i] - '0')
		a[v]++
	}
	a[x]++
	ret := 0

	c := []int{1, 10, 100, 1000, 10000, 100000}
	for i := 1; i <= 9; i++ {
		ret += i * c[a[i]]
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	K := getI()
	s := getS()
	t := getS()

	k := make([]int, 10)
	for i := 1; i < 10; i++ {
		k[i] = K
	}
	for _, e := range s {
		if e != '#' {
			k[int(e-'0')]--
		}
	}
	for _, e := range t {
		if e != '#' {
			k[int(e-'0')]--
		}
	}

	cnt := 0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i == j && k[i] < 2 {
				continue
			}
			if k[i] < 1 || k[j] < 1 {
				continue
			}
			S := calcScore(s, i)
			T := calcScore(t, j)
			if S > T {
				if i == j {
					cnt += k[i] * (k[i] - 1)
				} else {
					cnt += k[i] * k[j]
				}
			}
		}
	}

	tot := 9*K - 8
	tot = tot * (tot - 1)
	out(float64(cnt) / float64(tot))
}
