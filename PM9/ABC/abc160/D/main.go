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

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, x, y := getI(), getI()-1, getI()-1
	node := make([][]int, N)
	for i := 0; i < N-1; i++ {
		node[i] = append(node[i], i+1)
	}
	for i := N - 1; i > 0; i-- {
		node[i] = append(node[i], i-1)
	}
	node[x] = append(node[x], y)
	node[y] = append(node[y], x)

	k := make([]int, N)

	for i := 0; i < N; i++ {
		dist := make([]int, N)
		for j := 0; j < N; j++ {
			dist[j] = inf
		}
		q := []int{i}
		dist[i] = 0
		for len(q) != 0 {
			cur := q[0]
			q = q[1:]
			for _, nxt := range node[cur] {
				if dist[nxt] > dist[cur]+1 {
					dist[nxt] = dist[cur] + 1
					q = append(q, nxt)
				}
			}
		}
		for j := i + 1; j < N; j++ {
			if dist[j] != inf {
				k[dist[j]]++
			}
		}
	}

	for i := 1; i < N; i++ {
		out(k[i])
	}
}
