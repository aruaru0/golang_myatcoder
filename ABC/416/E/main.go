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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	const inf = int(1e18)
	n++
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = inf
		}
		d[i][i] = 0
	}

	for i := 0; i < m; i++ {
		a, b, c := getI(), getI(), getI()
		chmin(&d[a][b], c)
		chmin(&d[b][a], c)
	}

	k, T := getI(), getI()
	for i := 0; i < k; i++ {
		a := getI()
		chmin(&d[a][0], T)
		chmin(&d[0][a], 0)
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				chmin(&d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	ans := 0
	add := func(a, b, sign int) {
		if a == 0 || b == 0 {
			return
		}
		if d[a][b] != inf {
			ans += d[a][b] * sign
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			add(i, j, 1)
		}
	}

	update := func(a, b, c int) {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				add(i, j, -1)
				chmin(&d[i][j], d[i][a]+c+d[b][j])
				add(i, j, 1)
			}
		}
	}

	q := getI()
	for qi := 0; qi < q; qi++ {
		t := getI()
		switch t {
		case 1:
			a, b, c := getI(), getI(), getI()
			update(a, b, c)
			update(b, a, c)
		case 2:
			a := getI()
			update(a, 0, T)
			update(0, a, 0)
		case 3:
			out(ans)
		}

	}

}
