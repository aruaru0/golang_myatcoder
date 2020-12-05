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

var N, M, K int
var s [][]byte

func check(n, y, x int) bool {

	a := make([]int, 10)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := int(s[y+i][x+j] - '0')
			a[v]++
		}
	}

	rect := n * n
	ok := false
	for i := 0; i < 10; i++ {
		if a[i]+K >= rect {
			ok = true
		}
	}
	return ok
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, K = getI(), getI(), getI()
	s = make([][]byte, N)
	for i := 0; i < N; i++ {
		s[i] = []byte(getS())
	}

	n := min(N, M)

	ans := 0
	for k := 1; k <= n; k++ {
		for i := 0; i <= N-k; i++ {
			for j := 0; j <= M-k; j++ {
				ok := check(k, i, j)
				if ok {
					ans = max(ans, k)
				}
			}
		}
	}
	out(ans)
}
