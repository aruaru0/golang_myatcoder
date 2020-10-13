package main

import (
	"bufio"
	"fmt"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := make([]string, 5)
	c := make([]int, 5)
	for i := 0; i < 5; i++ {
		s[i], c[i] = getS(), getI()
	}
	n := make([]int, 3)
	a := make([]map[string]int, 3)
	for i := 0; i < 3; i++ {
		n[i] = getI()
		a[i] = make(map[string]int, n[i])
		for j := 0; j < n[i]; j++ {
			a[i][getS()]++
		}
	}
	// out(s, c)
	// out(n, a)
	tot := n[0] * n[1] * n[2]
	pair := make([]int, 5)
	for i := 0; i < 5; i++ {
		pair[i] = 1
		for j := 0; j < 3; j++ {
			pair[i] *= a[j][s[i]]
		}
		pair[i] *= 5
	}
	// out(tot, pair)
	ex := 0
	for i := 0; i < 5; i++ {
		ex += pair[i] * c[i]
	}
	out(float64(ex) / float64(tot))
	for _, e := range pair {
		out(e)
	}
}
