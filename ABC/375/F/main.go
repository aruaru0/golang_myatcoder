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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, Q := getI(), getI(), getI()

	a, b, c := make([]int, M), make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		a[i], b[i], c[i] = getI()-1, getI()-1, getI()
	}

	t := make([]int, Q)
	x, y := make([]int, Q), make([]int, Q)
	m := make(map[int]bool)
	for i := 0; i < Q; i++ {
		t[i] = getI()
		if t[i] == 1 {
			x[i] = getI() - 1
			m[x[i]] = true
		} else {
			x[i], y[i] = getI()-1, getI()-1
		}
	}

	const inf = int(1e18)
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
	// 追加しないノードを登録する
	for i := 0; i < M; i++ {
		if m[i] {
			continue
		}
		d[a[i]][b[i]] = c[i]
		d[b[i]][a[i]] = c[i]
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

	ans := make([]int, 0)
	for i := Q - 1; i >= 0; i-- {
		if t[i] == 1 {
			idx := x[i]
			u, v, w := a[idx], b[idx], c[idx]
			// d[a[idx]][b[idx]] = c[idx]
			// d[b[idx]][a[idx]] = c[idx]
			for i := 0; i < N; i++ {
				for j := 0; j < N; j++ {
					if d[i][u]+w+d[v][j] < d[i][j] {
						d[i][j] = d[i][u] + w + d[v][j]
					}
					if d[i][v]+w+d[u][j] < d[i][j] {
						d[i][j] = d[i][v] + w + d[u][j]
					}
				}
			}
		} else {
			u, v := x[i], y[i]
			ans = append(ans, d[u][v])
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] == inf {
			ans[i] = -1
		}
		out(ans[i])
	}

}
