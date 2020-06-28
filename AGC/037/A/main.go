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

const inf = 100

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := getString()
	N := len(s)
	s += "@@@@@"

	dp := make([][3]int, N+10)
	for i := 0; i < N+1; i++ {
		for j := 0; j < 3; j++ {
			dp[i][j] = -inf
		}
	}
	dp[0][0] = 0

	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			for k := 1; k < 3; k++ {
				if dp[i][j] >= 0 {
					if j == k {
						if k == 1 {
							if s[i-1] == s[i] {
								continue
							}
						} else {
							if s[i-2] == s[i] && s[i-1] == s[i+1] {
								continue
							}
						}
					}
					dp[i+k][k] = max(dp[i+k][k], dp[i][j]+1)
				}
			}
		}

	}

	out(max(dp[N][1], dp[N][2]))
}
