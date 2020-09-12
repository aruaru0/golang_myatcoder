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

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

const inf = int(1e15)

func main() {
	sc.Split(bufio.ScanWords)
	A, B, T := getInt(), getInt(), getInt()

	if T%A == 0 || T%B == 0 {
		out(T)
		return
	}

	l := lcm(A, B)
	k := max(0, T/l-1)
	d := 0

	if k*l <= T {
		d = k * l
		T -= k * l
	}

	n := (T + B - 1) / B

	// out(T, k, l, k*l, n)

	ans := inf
	for i := 0; i <= n+1; i++ {
		rest := T - B*i
		m := (rest + A - 1) / A
		if m < 0 {
			m = 0
		}
		ans = min(ans, B*i+A*m)
		// out(i, m, ans)
	}
	out(ans + d)
}
