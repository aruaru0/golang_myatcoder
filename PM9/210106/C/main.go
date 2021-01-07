package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

const full = int(1e10)

var N int

func match(s, t string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != t[i%len(t)] {
			return false
		}
	}
	return true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	T := getS()

	// 110パターンになっているか
	f := ""
	if match(T, "110") {
		f = ""
	} else if match(T, "101") {
		f = "10"
	} else if match(T, "011") {
		f = "0"
	} else {
		out(0)
		return
	}

	if T == "1" {
		out(full * 2)
		return
	}
	// if T == "0" || T == "11" {
	// 	out(full)
	// 	return
	// }

	n := strings.Count(T, "110")
	t := T[n*3+len(f):]
	if f != "" {
		n++
	}
	if t != "" {
		n++
	}
	out(full - n + 1)
	//out(N, n, f, t)
}
