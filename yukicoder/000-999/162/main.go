package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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
	n int
	p float64
}

var P [3]float64
var ns [][]pair

func prob(prev, next int) float64 {
	res := 1.0
	for i := 0; i < 14; i++ {
		if ((prev >> i) & 1) == 0 {
			continue
		}
		var p float64
		if i == 0 {
			p = P[(prev>>1)&1]
		} else if i == 13 {
			p = P[(prev>>12)&1]
		} else {
			p = P[(prev>>(i-1)&1)+(prev>>(i+1))&1]
		}

		if next>>i&1 == 1 {
			res *= 1 - p
		} else {
			res *= p
		}
	}
	return res
}

var dp [21][1 << 14]float64

func main() {
	sc.Split(bufio.ScanWords)
	N := 80 - getInt()
	p := getInts(3)
	for i := 0; i < 3; i++ {
		P[i] = float64(p[i]) / 100.0
	}

	ns = make([][]pair, 1<<14)
	for i := 0; i < 1<<14; i++ {
		for j := 0; j < 1<<14; j++ {
			if i&j != j {
				continue
			}
			ns[i] = append(ns[i], pair{j, prob(i, j)})
		}
	}

	dp[0][(1<<14)-1] = 1.0
	for i := 0; i < N; i++ {
		for j := 0; j < 1<<14; j++ {
			for _, p := range ns[j] {
				dp[i+1][p.n] += dp[i][j] * p.p
			}
		}
	}

	res := 0.0
	for i := 0; i < 1<<14; i++ {
		res += dp[N][i] * float64(bits.OnesCount(uint(i)))
	}
	out(res * 2)
}
