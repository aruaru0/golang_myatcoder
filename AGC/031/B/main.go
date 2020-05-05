package main

import (
	"bufio"
	"fmt"
	"os"
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

// Data :
type Data struct {
	f, t int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].f < p[j].f
}

func lowerBound(a []int, x int) int {
	l := 0
	r := len(a) - 1
	for l <= r {
		m := (l + r) / 2
		if a[m] >= x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(a []int, x int) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m] <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

const maxN = 200100
const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	c := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = getInt()
	}

	pos := make([][]int, maxN)
	for i := 0; i < N; i++ {
		pos[c[i]] = append(pos[c[i]], i)
	}

	// out(c)
	// out(pos[:10])

	dp := make([]int, N+1)
	dp[0] = 1
	for i := 1; i <= N; i++ {
		dp[i] += dp[i-1]
		dp[i] %= mod

		col := c[i-1]
		idx := lowerBound(pos[col], i-1)
		// out(col, pos[col], idx)
		if idx > 0 {
			j := pos[col][idx-1]
			if (i-1)-j > 1 {
				// out("add", i, j, "val", dp[j+1])
				dp[i] += dp[j+1]
				dp[i] %= mod
			}
		}
		// out(dp[1 : N+1])
	}
	out(dp[N])
}
