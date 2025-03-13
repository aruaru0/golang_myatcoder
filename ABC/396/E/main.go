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

type pair struct {
	to, xor int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	node := make([][]pair, n)
	for i := 0; i < m; i++ {
		a, b, c := getI()-1, getI()-1, getI()
		node[a] = append(node[a], pair{b, c})
		node[b] = append(node[b], pair{a, c})
	}

	var dfs func(v, c int)

	ok := true
	ans := make([]int, n)
	for k := 0; k < 30; k++ {
		col := make([]int, n)
		for i := 0; i < n; i++ {
			col[i] = -1
		}

		for sv := 0; sv < n; sv++ {
			if col[sv] == -1 {
				vs := make([][]int, 2)
				dfs = func(v, c int) {
					if col[v] != -1 {
						if col[v] != c {
							ok = false
						}
						return
					}
					col[v] = c
					vs[c] = append(vs[c], v)
					for _, e := range node[v] {
						dfs(e.to, c^((e.xor>>k)&1))
					}
				}
				dfs(sv, 0)
				if len(vs[0]) < len(vs[1]) {
					vs[0], vs[1] = vs[1], vs[0]
				}
				for _, e := range vs[1] {
					ans[e] |= 1 << k
				}
			}
		}
	}
	if !ok {
		out(-1)
	} else {
		for i := 0; i < n; i++ {
			fmt.Fprint(wr, ans[i], " ")
		}
		out()
	}

}
