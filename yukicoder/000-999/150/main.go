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

var good = "good"
var problem = "problem"

func solve(s string) {
	ans := len(s)
	n := len(s) - len(good) - len(problem) + 1
	for i := 0; i < n; i++ {
		cnt0 := 0
		for j := 0; j < len(good); j++ {
			if s[i+j] != good[j] {
				cnt0++
			}
		}
		// out(cnt0)
		m := len(s) - len(problem) + 1
		// out(i+4, m)
		for j := i + 4; j < m; j++ {
			cnt1 := 0
			for k := 0; k < len(problem); k++ {
				if s[j+k] != problem[k] {
					cnt1++
				}
			}
			ans = min(ans, cnt0+cnt1)
		}
	}
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	T := getInt()
	for i := 0; i < T; i++ {
		s := getString()
		solve(s)
	}
}
