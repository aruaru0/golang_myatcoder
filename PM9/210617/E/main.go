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

func check(s string) bool {
	if s[:3] == "AGC" || s[1:] == "AGC" {
		return true
	}
	return false
}

func f(s string, c byte) (string, bool) {
	ret := s + string(c)
	if check(ret) {
		return "", false
	}
	tmp := make([]byte, 4)
	tmp[0], tmp[1], tmp[2], tmp[3] = ret[1], ret[0], ret[2], ret[3]
	if check(string(tmp)) {
		return "", false
	}
	tmp[0], tmp[1], tmp[2], tmp[3] = ret[0], ret[2], ret[1], ret[3]
	if check(string(tmp)) {
		return "", false
	}
	tmp[0], tmp[1], tmp[2], tmp[3] = ret[0], ret[1], ret[3], ret[2]
	if check(string(tmp)) {
		return "", false
	}
	return ret[1:], true
}

const mod = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	dp := make(map[string]int)
	dp["xxx"] = 1

	s := "ACGT"
	for i := 0; i < N; i++ {
		tmp := make(map[string]int)
		for e, i := range dp {
			for _, c := range s {
				next, ok := f(e, byte(c))
				if ok {
					tmp[next] += i
					tmp[next] %= mod
				}
			}
		}
		dp = tmp
	}

	tot := 0
	for _, e := range dp {
		tot += e
		tot %= mod
	}
	out(tot)
}
