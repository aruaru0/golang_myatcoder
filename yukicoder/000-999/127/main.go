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
	A := getInts(N)

	var dpL [3030][3030]int
	var dpR [3030][3030]int

	for i := 0; i < N; i++ {
		dpL[i][i] = 1
		dpR[i][i] = 1
	}
	for n := 1; n < N; n++ {
		for i := 0; i < N-n; i++ {
			j := i + n
			dpL[i][j] = dpL[i][j-1]
			dpR[i][j] = dpR[i+1][j]
			if A[i] < A[j] {
				dpL[i][j] = max(dpL[i][j], dpR[i+1][j]+1)
			}
			if A[i] > A[j] {
				dpR[i][j] = max(dpR[i][j], dpL[i][j-1]+1)
			}
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			ans = max(ans, max(dpL[i][j], dpR[i][j]))
		}
	}
	out(ans)
}
