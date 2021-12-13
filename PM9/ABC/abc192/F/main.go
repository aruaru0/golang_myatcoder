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

const inf = int(math.MaxInt64)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, X := getI(), getI()
	A := getInts(N)

	var dp [101][101][101]int
	ans := inf
	for k := 1; k < N+1; k++ {
		for i := 0; i < N+1; i++ {
			for use := 0; use < k+1; use++ {
				for mo := 0; mo < k; mo++ {
					dp[i][use][mo] = inf
				}
			}
		}
		dp[0][0][X%k] = X
		for i := 0; i < N; i++ {
			for use := 0; use < k+1; use++ {
				for mo := 0; mo < k; mo++ {
					if dp[i][use][mo] != inf {
						dp[i+1][use][mo] = min(dp[i+1][use][mo], dp[i][use][mo])
						if use < k {
							dp[i+1][use+1][(((mo-A[i])%k)+k)%k] = min(dp[i+1][use+1][(((mo-A[i])%k)+k)%k], dp[i][use][mo]-A[i])
						}
					}
				}
			}
		}
		if dp[N][k][0] != inf {
			ans = min(ans, dp[N][k][0]/k)
		}
	}
	out(ans)
}
