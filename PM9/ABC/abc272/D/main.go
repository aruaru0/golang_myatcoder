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

type pair struct {
	x, y int
}

func pos2idx(x, y, N int) int {
	return y*N + x
}

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	d := make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([]int, N)
	}

	m := make([]pair, 0)
	for i := -N; i < N; i++ {
		for j := -N; j < N; j++ {
			if i*i+j*j == M {
				m = append(m, pair{i, j})
			}
		}
	}

	node := make([][]int, N*N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for _, e := range m {
				nx, ny := j+e.x, i+e.y
				if nx < 0 || nx >= N || ny < 0 || ny >= N {
					continue
				}
				node[pos2idx(i, j, N)] = append(node[pos2idx(i, j, N)], pos2idx(nx, ny, N))
				// node[pos2idx(nx, ny, N)] = append(node[pos2idx(nx, ny, N)], pos2idx(i, j, N))
			}
		}
	}

	used := make([]bool, N*N)
	dist := make([]int, N*N)
	for i := 0; i < N*N; i++ {
		dist[i] = inf
	}
	q := []int{0}
	used[0] = true
	dist[0] = 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for _, e := range node[cur] {
			if used[e] {
				continue
			}
			if dist[e] > dist[cur]+1 {
				dist[e] = dist[cur] + 1
				q = append(q, e)
				used[e] = true
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			v := dist[pos2idx(i, j, N)]
			if v == inf {
				v = -1
			}
			fmt.Fprint(wr, v, " ")
		}
		out()
	}
}
