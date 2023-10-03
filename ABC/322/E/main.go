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

func pos(k, bit int) []int {
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = (bit >> (i * 3)) & 7
	}
	return ret
}

func ipos(pos []int) int {
	ret := 0
	for i := 0; i < len(pos); i++ {
		ret += pos[i] << (i * 3)
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K, P := getI(), getI(), getI()

	c := make([]int, N)
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		c[i] = getI()
		a[i] = getInts(K)
	}

	const inf = int(1e18)
	// dp[bit] (bit>>(3*k)) k個目のパラメータが(bit>>(3*k) % 7)である場合の最小値
	bit := 1 << (K * 3)
	dp := make([]int, bit)
	for i := 0; i < bit; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < N; i++ {
		tmp := make([]int, bit)
		for j := 0; j < bit; j++ {
			tmp[j] = inf
		}
		for j := 0; j < bit; j++ {
			// cur := pos(K, j)
			nxt := pos(K, j)
			ok := true
			for k := 0; k < K; k++ {
				if nxt[k]+a[i][k] > P {
					nxt[k] = P
				} else {
					nxt[k] += a[i][k]
				}
			}
			tmp[j] = min(tmp[j], dp[j])
			if ok {
				idx := ipos(nxt)
				tmp[idx] = min(tmp[idx], dp[j]+c[i])
			}
		}
		dp = tmp
		// out(tmp)
	}

	p := make([]int, K)
	for i := 0; i < K; i++ {
		p[i] = P
	}
	idx := ipos(p)
	if dp[idx] == inf {
		out(-1)
		return
	}
	out(dp[idx])

}
