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

var used []bool
var node [][]int

func dfs(cur, prev int) bool {
	used[cur] = true
	ret := false
	for _, e := range node[cur] {
		if e == prev {
			continue
		}
		if used[e] == true {
			return true
		}
		ret = ret || dfs(e, cur)
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	node = make([][]int, N*2)
	red, blue := 0, N
	for i := 0; i < N; i++ {
		node[red+i] = append(node[red+i], blue+i)
		node[blue+i] = append(node[blue+i], red+i)
	}

	for i := 0; i < M; i++ {
		a, b, c, d := getI()-1, getS(), getI()-1, getS()
		if b[0] == 'B' {
			a += blue
		}
		if d[0] == 'B' {
			c += blue
		}
		node[a] = append(node[a], c)
		node[c] = append(node[c], a)
	}

	x, y := 0, 0
	used = make([]bool, N*2)
	for i := 0; i < 2*N; i++ {
		if used[i] == false {
			ret := dfs(i, -1)
			if ret == true {
				y++
			} else {
				x++
			}
		}

	}
	out(y, x)
}
