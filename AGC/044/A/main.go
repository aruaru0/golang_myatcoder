package main

import (
	"bufio"
	"fmt"
	"math"
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

const inf = math.MaxInt64
const size = 100000000

var a, b, c, d int
var dp map[int]int

func solve(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return d
	}
	// out(n)
	if dp[n] != 0 {
		return dp[n]
	}
	// out(n)
	var r, x int
	r = n % 5
	x = n - r
	ret := solve(x/5) + c + d*r
	if x != n && x != 0 {
		x += 5
		ret = min(ret, solve(x/5)+c+(5-r)*d)
	}

	r = n % 3
	x = n - r
	ret = min(ret, solve(x/3)+b+d*r)
	if x != n && x != 0 {
		x += 3
		ret = min(ret, solve(x/3)+b+(3-r)*d)
	}
	r = n % 2
	x = n - r
	ret = min(ret, solve(x/2)+a+d*r)
	if x != n && x != 0 {
		x += 2
		ret = min(ret, solve(x/2)+a+(2-r)*d)
	}

	if n*d/d == n {
		ret = min(ret, n*d)
	}
	dp[n] = ret
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	T := getInt()
	for i := 0; i < T; i++ {
		n := getInt()
		a, b, c, d = getInt(), getInt(), getInt(), getInt()
		dp = make(map[int]int)
		out(solve(n))
	}
}
