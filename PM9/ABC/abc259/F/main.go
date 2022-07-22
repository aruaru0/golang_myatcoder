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

type edge struct {
	to, cost int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	d := getInts(n)
	node := make([][]edge, n)
	for i := 0; i < n-1; i++ {
		a, b, c := getI()-1, getI()-1, getI()
		node[a] = append(node[a], edge{b, c})
		node[b] = append(node[b], edge{a, c})
	}

	const inf int = 1e18

	var dfs func(cur, prev int) []int
	dfs = func(cur, prev int) []int {
		e := []int{0}
		base := 0
		for _, v := range node[cur] {
			if v.to == prev {
				continue
			}
			r := dfs(v.to, cur)
			r[1] += max(0, v.cost)
			e = append(e, max(r[1]-r[0], 0))
			base += r[0]
		}
		res := []int{base, base}
		sort.Slice(e, func(i, j int) bool {
			return e[i] > e[j]
		})
		for i := 0; i < d[cur]; i++ {
			res[0] += e[i]
		}
		for i := 0; i < d[cur]-1; i++ {
			res[1] += e[i]
		}
		if d[cur] == 0 {
			res[1] = -inf
		}
		return res
	}

	dp := dfs(0, -1)
	out(dp[0])
}
