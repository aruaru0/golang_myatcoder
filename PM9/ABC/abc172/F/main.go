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

	// 解答を確認しながらgoに移植
	n := getI()
	_a := getInts(n)
	x := 0
	for i := 2; i < n; i++ {
		x ^= _a[i]
	}
	a, b := _a[0], _a[1]
	s := a + b

	K := 43
	dp := make([][2][2]int, K)
	for i := 0; i < K; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	dp[0][0][0] = 0
	v := 1
	for i := 0; i < K-1; i++ {
		cx, cs, ca := x&1, s&1, a&1
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				if dp[i][j][k] == -1 {
					continue
				}
				for na := 0; na < 2; na++ {
					for nb := 0; nb < 2; nb++ {
						ni, nj, nk := i+1, 0, k
						if (na ^ nb) != cx {
							continue
						}
						ns := na + nb + j
						if ns%2 != cs {
							continue
						}
						if ns >= 2 {
							nj = 1
						}
						if ca < na {
							nk = 1
						} else if ca == na {
							nk = k
						} else {
							nk = 0
						}
						chmax(&dp[ni][nj][nk], dp[i][j][k]|(v*na))
					}
				}
			}
		}
		x >>= 1
		s >>= 1
		a >>= 1
		v <<= 1
	}

	a = dp[K-1][0][0]
	if a == -1 || a == 0 {
		out(-1)
		return
	}
	ans := _a[0] - a
	out(ans)
}
