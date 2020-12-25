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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	var a [35][31][31]int
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			a[0][i][j] = getI()
		}
	}

	for n := 0; n < 34; n++ {
		for k := 0; k < M; k++ {
			for i := 0; i < M; i++ {
				for j := 0; j < M; j++ {
					a[n+1][i][j] = max(a[n+1][i][j], a[n][i][k]+a[n][k][j])
				}
			}
		}
	}

	N--
	ans := 0
	var b [35][31][31]int
	for n := 0; n < 34; n++ {
		if (N>>n)&1 != 0 {
			for k := 0; k < M; k++ {
				for i := 0; i < M; i++ {
					for j := 0; j < M; j++ {
						b[n+1][i][j] = max(b[n+1][i][j], b[n][i][k]+a[n][k][j])
						ans = max(ans, b[n+1][i][j])
					}
				}
			}
		} else {
			for i := 0; i < M; i++ {
				for j := 0; j < M; j++ {
					b[n+1][i][j] = b[n][i][j]
				}
			}
		}
	}
	out(ans)
}
