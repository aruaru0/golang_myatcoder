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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, L := getI(), getI(), getI()
	d := make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			d[i][j] = inf
		}
	}
	for i := 0; i < M; i++ {
		a, b, c := getI()-1, getI()-1, getI()
		d[a][b] = c
		d[b][a] = c
	}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if d[i][j] > d[i][k]+d[k][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if d[i][j] <= L {
				d[i][j] = 1
			} else {
				d[i][j] = inf
			}
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if d[i][j] > d[i][k]+d[k][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}

	Q := getI()
	for i := 0; i < Q; i++ {
		s, t := getI()-1, getI()-1
		dist := d[s][t]
		if dist == inf {
			out(-1)
			continue
		}
		out(dist - 1)
	}

}
