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
	u, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	node := make([][]int, N)
	m := make(map[pair]bool)
	for i := 0; i < N-1; i++ {
		u, v := getI()-1, getI()-1
		if u > v {
			u, v = v, u
		}
		node[u] = append(node[u], v)
		node[v] = append(node[v], u)
		m[pair{u, v}] = true
	}

	p := make([][]int, 2)
	var dfs func(cur, pre, flg int)
	dfs = func(cur, pre, flg int) {
		p[flg] = append(p[flg], cur)
		for _, e := range node[cur] {
			if e == pre {
				continue
			}
			dfs(e, cur, 1^flg)
		}
	}

	dfs(0, -1, 0)

	rest := make(map[pair]bool, 0)
	for _, x := range p[0] {
		for _, y := range p[1] {
			u, v := x, y
			if u > v {
				u, v = v, u
			}
			if !m[pair{u, v}] {
				rest[pair{u, v}] = true
			}
		}
	}
	turn := 0
	if len(rest)%2 == 1 {
		fmt.Println("First")
		turn = 0
	} else {
		fmt.Println("Second")
		turn = 1
	}

	for {
		if turn == 0 {
			u, v := 0, 0
			for k := range rest {
				u, v = k.u, k.v
				break
			}
			fmt.Println(u+1, v+1)
			delete(rest, pair{u, v})
		} else {
			u, v := getI(), getI()
			if u == -1 {
				return
			}
			if u > v {
				u, v = v, u
			}
			delete(rest, pair{u - 1, v - 1})
		}
		turn ^= 1
	}
}
