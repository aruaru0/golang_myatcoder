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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	s := getString()
	digit := 1
	n := 0
	cnt := 0
	sum := 0
	m := make([][]int, 0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '?' {
			b := make([]int, 10)
			for j := 0; j < 10; j++ {
				b[j] = (j * digit) % 13
				// out(j, digit, b[j])
			}
			sort.Ints(b)
			m = append(m, b)
			// out(b)
			cnt++
		} else {
			x := int(s[i] - '0')
			n = x * digit
			n %= 13
			sum += n
			sum %= 13
			// out(x, n)
		}
		digit *= 10
		digit %= 13
	}
	// out(sum)
	// for i := 0; i < cnt; i++ {
	// 	out(m[i])
	// }

	dp := make([][]int, cnt+1)
	for i := 0; i <= cnt; i++ {
		dp[i] = make([]int, 13)
	}

	dp[0][0] = 1
	for i := 0; i < cnt; i++ {
		for j := 0; j < 10; j++ {
			// out(i, j, m[i][j])
			for k := 0; k < 13; k++ {
				next := (m[i][j] + k) % 13
				dp[i+1][next] += dp[i][k]
				dp[i+1][next] %= mod
				// out("-->", (m[i][j]+k)%13)
			}
			// dp[i+1][xx] S+= dp[i-1][yy]
		}
	}

	pos := []int{5, 4, 3, 2, 1, 0, 12, 11, 10, 9, 8, 7, 6}
	out(dp[cnt][pos[sum]])
}
