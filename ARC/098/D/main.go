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

func check(a, b int) bool {
	for i := 0; i <= 32; i++ {
		x := (a >> uint(i)) % 2
		y := (b >> uint(i)) % 2
		if x == 1 && y == 1 {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	cur := 0
	ans := 0
	l := 0

	for i := 0; i < N; i++ {
		f := check(cur, a[i])
		// out(f)
		if f {
			cur ^= a[i]
		} else {
			for j := 1; j <= i-l; j++ {
				ans += j
			}
			for l < i {
				cur ^= a[l]
				l++
				f := check(cur, a[i])
				if f {
					cur ^= a[i]
					break
				}
			}
			for j := 1; j <= i-l; j++ {
				ans -= j
			}
			// out("ans--", ans, l, i)
		}
	}

	// out("cnt", l)
	if l != N {
		for i := 1; i <= N-l; i++ {
			ans += i
		}
	}
	out(ans)
}
