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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	sc.Split(bufio.ScanWords)
	A, B := getInt(), getInt()

	dp := make([]int, 11000)
	dp[0] = 1
	for i := 0; i < 11000; i++ {
		if i+A < 11000 {
			dp[i+A] += dp[i]
		}
	}
	for i := 0; i < 11000; i++ {
		if i+B < 11000 {
			dp[i+B] += dp[i]
		}
	}

	cnt := 0
	flg := false
	for i := 0; i < 11000; i++ {
		if dp[i] == 0 {
			if A*B >= i {
				cnt++
			} else {
				flg = true
			}
		}
	}
	if flg == true {
		out(-1)
		return
	}
	out(cnt)
}
