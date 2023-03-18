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
	N, M := getI(), getI()
	a := getInts(N)
	x := make([]int, M)
	y := make([]int, M)
	z := make([]int, M)
	for i := 0; i < M; i++ {
		x[i], y[i], z[i] = getI()-1, getI()-1, getI()-1
	}

	// スタート状態をビット列に変換
	p := 0
	for i := len(a) - 1; i >= 0; i-- {
		p <<= 1
		p |= a[i]
	}

	// 状態をビット列として捉え、ｘ、ｙ、ｚの操作で、変換されるビット列まで辺を引く
	G := make([][]int, 1<<N)
	for i := 0; i < M; i++ {
		q := 1<<x[i] | 1<<y[i] | 1<<z[i]
		for j := 0; j < 1<<N; j++ {
			G[j] = append(G[j], j^q)
		}
	}

	// BSF
	const inf = int(1e18)
	dist := make([]int, 1<<N)
	for i := 0; i < 1<<N; i++ {
		dist[i] = inf
	}
	dist[p] = 0
	deq := []int{p}
	for len(deq) != 0 {
		cur := deq[0]
		deq = deq[1:]
		for _, nxt := range G[cur] {
			if dist[nxt] > dist[cur]+1 {
				dist[nxt] = dist[cur] + 1
				deq = append(deq, nxt)
			}
		}
	}

	// 全て１の状態の値が結果
	ans := dist[(1<<N)-1]
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}
