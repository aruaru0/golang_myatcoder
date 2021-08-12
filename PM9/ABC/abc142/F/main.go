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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	A := make([]int, M)
	B := make([]int, M)
	node := make([][]int, N)
	for i := 0; i < M; i++ {
		A[i], B[i] = getI()-1, getI()-1
		node[A[i]] = append(node[A[i]], B[i])
	}

	mi := inf
	ans := make([]int, 0)
	for er := 0; er < M; er++ {
		// 辺の始点と終点を選ぶ
		st := A[er]
		gr := B[er]

		vis := make([]bool, N)
		dist := make([]int, N)
		back := make([]int, N)

		// BSF
		q := []int{gr}
		vis[gr] = true
		dist[gr] = 0
		for len(q) != 0 {
			cu := q[0]
			q = q[1:]
			for _, to := range node[cu] {
				if vis[to] {
					continue
				}
				// 訪問フラグ
				vis[to] = true
				// 距離
				dist[to] = dist[cu] + 1
				// 来たノード
				back[to] = cu
				q = append(q, to)
			}
		}
		// 始点に到着していない場合はcontinue
		if !vis[st] {
			continue
		}

		// 距離は始点までの距離＋１辺
		l := dist[st] + 1
		if mi > l { // 最小値を更新
			mi = l
			ans = make([]int, 0)
			// st -> ... -> grをたどる
			x := st
			for x != gr {
				ans = append(ans, x)
				x = back[x]
			}
			ans = append(ans, gr)
		}
	}
	if mi == inf {
		out(-1)
		return
	}
	sort.Ints(ans)
	out(len(ans))
	for _, e := range ans {
		out(e + 1)
	}
}
