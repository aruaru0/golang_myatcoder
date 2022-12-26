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

func pow(x, n int) int {
	ans := 1
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x
		}
		x = x * x
		n = n >> 1
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N+1)
	sum := 0
	for i := 0; i <= N; i++ {
		a[i] = getInt()
		sum += a[i]
	}

	ans := 0
	x := 1
	for i := 0; i <= N; i++ {
		ans += x
		if x < a[i] {
			out(-1)
			return
		}
		x -= a[i]
		sum -= a[i]
		x *= 2
		x = min(x, sum)
	}
	out(ans)
}
