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

// 思いつかなくてコード写経。なるほどとは思うけど思いつかない
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 10000000)
	s := getString()
	n := len(s)

	dpi := make([]int, 26)
	dp := make([][26]int64, n+1)
	var ans int64
	for i := 0; i < n; i++ {
		ch := int(s[i] - 'a')
		for j := 0; j < 26; j++ {
			dp[i+1][j] = dp[i][j]

			if ch == j {
				dp[i+1][j] += int64(i + 1 - dpi[j])
				dpi[j] = i + 1
			}
			ans += dp[i+1][j]
		}
	}
	avg := float64(ans) / float64(int64(n)*int64(n+1)/2)
	out(avg)
}
