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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node := make([][]int, N)
	for i := 0; i < N-1; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
	}
	c := getInts(N)
	sum := 0
	for i := 0; i < N; i++ {
		sum += c[i]
	}

	x := -1
	{
		// 重心を求める
		var dfs func(int, int) int
		dfs = func(cur, prev int) int {
			ret := c[cur]
			mx := 0
			for _, e := range node[cur] {
				if e == prev {
					continue
				}
				now := dfs(e, cur)
				mx = max(mx, now) // 一番大きな重みの部分木
				ret += now        // この頂点の部分木の重み
			}
			mx = max(mx, sum-ret) // 部分木の重みと、それ以外の部分の重みの大きな方
			if mx*2 <= sum {
				x = cur
			}
			return ret
		}
		dfs(0, -1)
	}

	// 根の頂点が決まったらあとは計算するだけ
	ans := 0
	var dfs func(int, int, int)
	dfs = func(cur, prev, dist int) {
		ans += dist * c[cur]
		for _, e := range node[cur] {
			if e == prev {
				continue
			}
			dfs(e, cur, dist+1)
		}
	}
	dfs(x, -1, 0)
	out(ans)
}
