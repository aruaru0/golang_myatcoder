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

const maxN = 55
const mod = 1234567891

var A [maxN]int
var dp [50500]int

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, M := getInt(), getInt()
	A := getInts(N)

	dp[0] = 1
	for M != 0 {
		for i := 0; i < N; i++ {
			for j := 50000; j >= A[i]; j-- {
				dp[j] += dp[j-A[i]]
				dp[j] %= mod
			}
		}
		for i := 0; i < 50*500; i++ {
			dp[i] = dp[i*2+M%2]
		}
		for i := 0; i < 50*500; i++ {
			dp[i+50*500] = 0
		}
		M /= 2
	}
	out(dp[0])
}
