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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type Edge struct {
	to, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	a := make([]int, N-1)
	b := make([]int, N-1)
	node := make([][]Edge, N)
	for i := 0; i < N-1; i++ {
		a[i], b[i] = getI()-1, getI()-1
		node[a[i]] = append(node[a[i]], Edge{b[i], i})
		node[b[i]] = append(node[b[i]], Edge{a[i], i})

	}

	path := make([]int, 0)
	edges := make([][]int, N-1)

	var dfs func(cur, prev int)
	dfs = func(cur, prev int) {
		path = append(path, cur)

		for _, e := range node[cur] {
			if e.to == prev {
				continue
			}
			edges[e.idx] = append(edges[e.idx], len(path))
			dfs(e.to, cur)
			edges[e.idx] = append(edges[e.idx], len(path))
		}

		path = append(path, cur)
	}

	dfs(0, -1)

	// out(path)
	// out(edges)

	cnt := make([]int, len(path)+1)
	for i := 0; i < N-1; i++ {
		l, r := edges[i][0], edges[i][1]
		// out(l, r, r-l, len(path)-(r-l))
		x, y := r-l, len(path)-(r-l)
		if x > y {
			cnt[l] += i + 1
			cnt[r] -= i + 1
		} else {
			cnt[0] += i + 1
			cnt[len(path)] -= i + 1
			cnt[l] -= i + 1
			cnt[r] += i + 1
		}
	}

	for i := 0; i < len(path); i++ {
		cnt[i+1] += cnt[i]
	}

	ans := make([]int, N)
	for i := 0; i < len(path); i++ {
		ans[path[i]] += cnt[i]
	}

	for i := 0; i < N; i++ {
		ans[i] /= 2
	}
	outSlice(ans)

}
