package main

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K := getInt(), getInt()
	var dp1 [11][100]float64
	var dp2 [11][100]float64
	dp1[0][0] = 1.0
	dp2[0][0] = 1.0

	for i := 1; i <= N; i++ {
		for j := 1; j <= 6*i; j++ {
			if i <= K {
				for k := 4; k <= 6; k++ {
					if j-k >= 0 {
						dp1[i][j] += dp1[i-1][j-k] / 3.0
					}
				}
			} else {
				for k := 1; k <= 6; k++ {
					if j-k >= 0 {
						dp1[i][j] += dp1[i-1][j-k] / 6.0
					}
				}
			}
		}
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= 6*i; j++ {
			for k := 1; k <= 6; k++ {
				if j-k >= 0 {
					dp2[i][j] += dp2[i-1][j-k] / 6.0
				}
			}
		}
	}
	ans := 0.0
	for sum1 := 1; sum1 <= N*6; sum1++ {
		for sum2 := 0; sum2 < sum1; sum2++ {
			ans += dp1[N][sum1] * dp2[N][sum2]
		}
	}
	out(ans)
}
