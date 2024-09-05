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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
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
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			d[i][j] = inf
		}
	}
	u := make([]int, M)
	v := make([]int, M)
	t := make([]int, M)
	for i := 0; i < M; i++ {
		u[i], v[i], t[i] = getI()-1, getI()-1, getI()
		chmin(&d[u[i]][v[i]], t[i])
		chmin(&d[v[i]][u[i]], t[i])
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				chmin(&d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	Q := getI()
	for qi := 0; qi < Q; qi++ {
		K := getI()
		a := make([]int, K)
		b := make([]int, K)
		for i := 0; i < K; i++ {
			a[i] = i
			b[i] = getI() - 1
		}
		ans := inf
		for { // 通過パターンを全て網羅（通過する橋のパターン）
			for bit := 0; bit < 1<<K; bit++ { // 1のときu->v 0のときv->u（方向）
				cur, cost := 0, 0
				for i := 0; i < K; i++ {
					idx := a[i]
					if bit>>i%2 == 1 {
						cost += d[cur][u[b[idx]]] + t[b[idx]]
						cur = v[b[idx]]
					} else {
						cost += d[cur][v[b[idx]]] + t[b[idx]]
						cur = u[b[idx]]
					}
				}
				cost += d[cur][N-1]
				ans = min(ans, cost)
			}

			if NextPermutation(sort.IntSlice(a)) == false {
				break
			}
		}
		out(ans)
	}

}
