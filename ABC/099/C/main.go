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

const inf = 1001001001

func main() {
	sc.Split(bufio.ScanWords)

	a := make([]int, 0)
	n := 1
	for n < 100000 {
		n *= 6
		a = append(a, n)
	}
	n = 1
	for n < 100000 {
		n *= 9
		a = append(a, n)
	}
	a = append(a, 1)
	sort.Ints(a)
	N := getInt()

	dp := make([]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = inf
	}
	dp[0] = 0
	for k := 0; k < 100; k++ {
		for i := 0; i < len(a); i++ {
			for j := 0; j <= N; j++ {
				if j+a[i] <= N {
					dp[j+a[i]] = min(dp[j+a[i]], dp[j]+1)
				}
			}
		}
		//		for i := 0; i <= N; i++ {
		//			fmt.Printf("%d:%d ", i, dp[i])
		//		}
		//		out()
	}
	out(dp[N])
}
