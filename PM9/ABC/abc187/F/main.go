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

const inf = int(255)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()

	// G[i]に辺がつながっているbit集合
	G := make([]int, N)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		G[a] |= 1 << b
		G[b] |= 1 << a
	}

	n := 1 << N
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}
	dp[0] = 1
	// クリークに１を立てる
	for i := 0; i < n; i++ {
		for j := 0; j < N; j++ {
			// j bit目が立っていない場合はcontinue
			if (i & (1 << j)) == 0 {
				continue
			}
			// j bit目を消す
			t := i ^ (1 << j)
			// infで無い　かつ　G[j]のt bitがすべて立っている
			if dp[t] == 1 && G[j]&t == t {
				dp[i] = 1
			}
			break
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j != 0; j = (j - 1) & i {
			dp[i] = min(dp[i], dp[j]+dp[i^j])
		}
	}
	out(dp[n-1])
}
