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
	u := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return u
}

func upperBound(a []int, x int) int {
	u := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return u
}

var node [][]int
var enter []int
var leave []int
var cnt int

func dfs(cur, prev int) {
	enter[cur] = cnt
	cnt++
	for _, e := range node[cur] {
		if e == prev {
			continue
		}
		dfs(e, cur)
	}
	leave[cur] = cnt
}

type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	b := new(BIT)
	b.v = make([]int, n)
	return b
}
func (b BIT) sum(a int) int {
	ret := 0
	for i := a + 1; i > 0; i -= i & -i {
		ret += b.v[i-1]
	}
	return ret
}
func (b BIT) rangeSum(x, y int) int {
	if y == 0 {
		return 0
	}
	y--
	if x == 0 {
		return b.sum(y)
	} else {
		return b.sum(y) - b.sum(x-1)
	}
}
func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

type pair struct {
	idx, val int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	c := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = getI() - 1
	}
	node = make([][]int, N)
	for i := 0; i < N-1; i++ {
		a, b := getI()-1, getI()-1
		node[a] = append(node[a], b)
		node[b] = append(node[b], a)
	}
	enter = make([]int, N)
	leave = make([]int, N)
	cnt = 0
	dfs(0, -1)

	colors := make([][]int, N)
	for i := 0; i < N; i++ {
		colors[c[i]] = append(colors[c[i]], i)
	}

	bit := newBIT(N)
	for i := 0; i < N; i++ {
		bit.add(i, 1)
	}

	for i := 0; i < N; i++ {
		ans := N * (N + 1) / 2
		col := colors[i]
		sort.Slice(col, func(i, j int) bool {
			return enter[col[i]] > enter[col[j]]
		})
		log := make([]pair, 0)
		for _, v := range col {
			cnt := 1
			for _, u := range node[v] {
				if enter[u] < enter[v] {
					continue // 親ならコンテニュー
				}
				num := bit.rangeSum(enter[u], leave[u])
				ans -= num * (num + 1) / 2
				cnt += num
			}
			bit.add(enter[v], -cnt)
			log = append(log, pair{enter[v], cnt})
		}
		num := bit.rangeSum(0, N)
		ans -= num * (num + 1) / 2
		for _, e := range log {
			bit.add(e.idx, e.val)
		}
		out(ans)
	}
}
