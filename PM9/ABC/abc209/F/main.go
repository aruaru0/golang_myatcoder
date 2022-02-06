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

const mod = int(1e9 + 7)

// 2回目トライしたが自力では解けなかった
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	H := getInts(N)
	// dp[i][pos] :  i番目までの木の伐採順番が決まっていて、
	//               最後の木（i番目の木）がpos番目で伐採
	//               されているような組合せ
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, N+1)
	}
	tot := make([]int, N+1)
	dp[1][1] = 1
	for i := 1; i < N; i++ {
		if H[i-1] <= H[i] {
			tot[i] = dp[i][i]
			for pos := i - 1; pos >= 1; pos-- {
				tot[pos] = tot[pos+1] + dp[i][pos]
				tot[pos] %= mod
			}
			for pos2 := 1; pos2 <= i; pos2++ {
				dp[i+1][pos2] += tot[pos2]
				dp[i+1][pos2] %= mod
			}
		}
		if H[i-1] >= H[i] {
			tot[1] = dp[i][1]
			for pos := 2; pos <= i; pos++ {
				tot[pos] = tot[pos-1] + dp[i][pos]
				tot[pos] %= mod
			}
			for pos2 := 2; pos2 <= i+1; pos2++ {
				dp[i+1][pos2] += tot[pos2-1]
				dp[i+1][pos2] %= mod
			}
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans += dp[N][i]
		ans %= mod
	}
	out(ans)
}
