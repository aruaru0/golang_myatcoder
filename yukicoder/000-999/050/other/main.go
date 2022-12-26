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
	N := getInt()
	a := getInts(N)
	M := getInt()
	b := getInts(M)
	sort.Ints(a)
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	n := 1 << N
	c := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < N; j++ {
			if (i>>j)%2 == 1 {
				c[i] += a[j]
			}
		}
	}

	dp := make([][]bool, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true

	for i := 1; i <= M; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if j|k != j {
					continue
				}
				p := j & k
				x := j ^ k
				if c[x] > b[i-1] {
					continue
				}
				// out(p, x, c[x], i, j)
				dp[i][j] = dp[i][j] || dp[i-1][p]
			}
		}
	}

	ans := -1
	for i := 0; i <= M; i++ {
		// out(dp[i][n-1])
		if dp[i][n-1] == true {
			ans = i
			break
		}
	}
	out(ans)
}
