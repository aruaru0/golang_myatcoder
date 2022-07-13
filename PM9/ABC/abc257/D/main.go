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

var N int
var x, y, P []int

func canMove(md, from, to int) bool {
	return P[from]*md >= abs(x[from]-x[to])+abs(y[from]-y[to])
}

func check(md int) bool {
	// 開始位置をすべて探索
	for start := 0; start < N; start++ {
		// BSF
		visited := make(map[int]bool)
		que := make([]int, 0)
		visited[start] = true
		que = append(que, start)
		for len(que) != 0 {
			cur := que[0]
			que = que[1:]

			for to := 0; to < N; to++ {
				if to == cur {
					continue
				}
				// mdでつながっているかどうかチェックし、つながっている
				// かつ訪問していないなら、キューに追加
				if canMove(md, cur, to) && visited[to] == false {
					visited[to] = true
					que = append(que, to)
				}
			}
		}
		if len(visited) == N {
			return true
		}
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	x = make([]int, N)
	y = make([]int, N)
	P = make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i], P[i] = getI(), getI(), getI()
	}

	// 二分探索
	ng, ok := -1, int(4e9)+100
	for ng+1 != ok {
		md := (ok + ng) / 2
		if check(md) {
			ok = md
		} else {
			ng = md
		}
	}
	out(ok)
}
