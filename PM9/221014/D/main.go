package main

import (
	"bufio"
	"fmt"
	"math"
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

func between(c, x, y byte) bool {
	return x <= c && c <= y
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	c := getS()
	n := len(c)

	// 最初と最後の連続'_'列の長さ
	headUS, tailUS := 0, 0
	for i := 0; i < n && c[i] == '_'; i++ {
		headUS++
	}
	for i := n - 1; i >= 0 && c[i] == '_'; i-- {
		tailUS++
	}
	if headUS+tailUS >= n {
		out(c)
		return
	}

	// 最初と最後の連続'_'列を除去した文字列
	s := c[headUS : n-tailUS]
	// m := len(s)
	t := ""

	// snake -> camel
	if words := strings.Split(s, "_"); len(words) > 1 {
		for i, word := range words {
			if len(word) == 0 || between(word[0], '0', '9') {
				out(c)
				return
			}
			for j := range word {
				if between(word[j], 'A', 'Z') {
					out(c)
					return
				}
			}
			if i == 0 {
				t += word
			} else {
				t += strings.ToUpper(string(word[0])) + word[1:]
			}
		}
		out(strings.Repeat("_", headUS) + t + strings.Repeat("_", tailUS))
		return
	}

	// camel -> snake
	if !between(s[0], 'a', 'z') {
		out(c)
		return
	}
	for i := range s {
		if between(s[i], 'A', 'Z') {
			t += "_" + strings.ToLower(string(s[i]))
		} else {
			t += string(s[i])
		}
	}
	out(strings.Repeat("_", headUS) + t + strings.Repeat("_", tailUS))
}
