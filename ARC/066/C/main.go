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
	N := getInt()
	a := make([]int, N)
	m := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
		m[a[i]]++
	}
	// out(m)
	ans := 1
	if N%2 == 1 {
		if m[0] != 1 {
			out(0)
			return
		}
		for i := 2; i <= N-1; i += 2 {
			if m[i] != 2 {
				out(0)
				return
			}
			ans *= 2
			ans %= mod
		}
	} else {
		for i := 1; i <= N-1; i += 2 {
			if m[i] != 2 {
				out(0)
				return
			}
			ans *= 2
			ans %= mod
		}
	}
	out(ans)
}
