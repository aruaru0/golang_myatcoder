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

var M int
var node [][]int

type pat struct {
	x   [9]int
	pos int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	M = getI()
	node = make([][]int, 9)
	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
	}

	v := pat{}
	for i := 0; i < 8; i++ {
		v.x[getI()-1] = i + 1
	}
	pos := 0
	for i := 0; i < 9; i++ {
		if v.x[i] == 0 {
			pos = i
			break
		}
	}
	v.pos = pos

	ans := pat{}
	for i := 0; i < 8; i++ {
		ans.x[i] = i + 1
	}
	ans.pos = 8
	ok := false

	const inf = int(1e18)
	dist := make(map[pat]int)
	used := make(map[pat]bool)
	used[v] = true
	dist[v] = 0
	if v == ans {
		ok = true
	}
	q := []pat{v}
	for len(q) > 0 {
		cur := q[0]
		// out("------------------")
		// out("q", q, node[cur.pos])
		// out("used", used)
		// out("dist", used)
		q = q[1:]
		for _, e := range node[cur.pos] {
			p := pat{}
			for i := 0; i < 9; i++ {
				p.x[i] = cur.x[i]
			}
			p.pos = e
			p.x[e], p.x[cur.pos] = p.x[cur.pos], p.x[e]
			if used[p] == false {
				used[p] = true
				dist[p] = dist[cur] + 1
				if p == ans {
					ok = true
				}
				q = append(q, p)
			}
		}
	}

	if ok {
		out(dist[ans])
	} else {
		out("-1")
	}
}
