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
	n := getI()
	c := append([]int{0}, getInts(n-1)...)
	a := append([]int{0}, getInts(n-1)...)

	node := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 1; j <= c[i]; j++ {
			from := i - j
			node[from] = append(node[from], i)
		}
	}

	rt := []int{0}
	for i := 0; i < n; i++ {
		if a[i] == 1 {
			rt = append(rt, i)
		}
	}

	var dist []int

	bsf := func(cur, target, cnt int) {
		q := []int{cur}
		for len(q) != 0 {
			from := q[0]
			q = q[1:]
			dist[cnt] = cnt
			for _, to := range node[from] {
				if dist[to] > dist[from]+1 {
					dist[to] = dist[from] + 1
					q = append(q, to)
					if to == target {
						return
					}
				}
			}
		}
	}

	const inf = int(1e8)

	tot := 0
	for i := 0; i < len(rt)-1; i++ {
		from, to := rt[i], rt[i+1]
		dist = make([]int, n)
		for j := 0; j < n; j++ {
			dist[j] = inf
		}
		dist[from] = 0
		bsf(from, to, 0)
		tot += dist[to]
	}
	out(tot)
}
