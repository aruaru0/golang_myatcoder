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

const inf = int(1e18)

type pair struct {
	x, y int
}

func solve() {
	n, m := getI(), getI()
	c := getInts(n)
	to := make([][]int, n)
	for i := 0; i < m; i++ {
		a, b := getI()-1, getI()-1
		to[a] = append(to[a], b)
		to[b] = append(to[b], a)
	}
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = inf
		}
	}

	q := make([]pair, 0)
	push := func(i, j, d int) {
		// 既に訪問していればパスする
		if dist[i][j] != inf {
			return
		}
		dist[i][j] = d
		q = append(q, pair{i, j})
	}

	// 以下は、bsfの探索と同じ
	// スタートをキューに入れる
	push(0, n-1, 0)
	for len(q) != 0 {
		// キューから取り出す
		a, b := q[0].x, q[0].y
		q = q[1:]
		d := dist[a][b]
		// 次に行ける場所に遷移
		// 　行けるのは、to[a]とto[b]の組み合わせの全て
		for _, na := range to[a] {
			for _, nb := range to[b] {
				if c[na] == c[nb] {
					continue
				}
				// 行ける場合は、距離に＋１する
				push(na, nb, d+1)
			}
		}
	}

	ans := dist[n-1][0]
	if ans == inf {
		ans = -1
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		solve()
	}
}
