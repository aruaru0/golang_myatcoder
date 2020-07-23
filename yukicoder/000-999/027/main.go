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

func calc(x, a, b, c int) int {
	A := x / a
	B := x / b
	C := x / c
	ret := 100000
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				y := i*a + j*b + k*c
				if y == x {
					// out(x, a, b, c, "---", i, j, k)
					ret = min(ret, i+j+k)
				}
			}
		}
	}
	return ret
}

var memo []int

func rec(n, a, b, c int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return 100000
	}
	if memo[n] >= 0 {
		return memo[n]
	}
	memo[n] = 1 + min(min(rec(n-a, a, b, c),
		rec(n-b, a, b, c)),
		rec(n-c, a, b, c))
	return memo[n]
}

func main() {
	sc.Split(bufio.ScanWords)
	v0, v1, v2, v3 := getInt(), getInt(), getInt(), getInt()
	ans := 100000
	for a := 1; a <= 30; a++ {
		for b := 1; b <= 30; b++ {
			for c := 1; c <= 30; c++ {
				// memo = make([]int, 100)
				// for i := 0; i < 100; i++ {
				// 	memo[i] = -1
				// }
				// ans = min(ans,
				// 	rec(v0, a, b, c)+
				// 		rec(v1, a, b, c)+
				// 		rec(v2, a, b, c)+
				// 		rec(v3, a, b, c))
				n := calc(v0, a, b, c)
				n += calc(v1, a, b, c)
				n += calc(v2, a, b, c)
				n += calc(v3, a, b, c)
				ans = min(ans, n)
			}
		}
	}
	out(ans)
}
