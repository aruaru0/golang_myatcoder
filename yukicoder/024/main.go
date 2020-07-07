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
	x := make([]int, 10)
	for i := 0; i < N; i++ {
		a, b, c, d, r := getInt(), getInt(), getInt(), getInt(), getString()
		// out(a, b, c, d)
		if r == "NO" {
			x[a] = -1
			x[b] = -1
			x[c] = -1
			x[d] = -1
		} else {
			for j := 0; j < 10; j++ {
				if j == a || j == b || j == c || j == d {
					continue
				}
				x[j] = -1
			}
		}
	}
	// out(x)
	ans := 0
	for i := 0; i < 10; i++ {
		if x[i] == 0 {
			// out(i)
			ans = i
		}
	}
	out(ans)
}
