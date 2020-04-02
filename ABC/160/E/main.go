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

func main() {
	sc.Split(bufio.ScanWords)

	X, Y, A, B, C := getInt(), getInt(), getInt(), getInt(), getInt()
	p := make([]int, A)
	q := make([]int, B)
	r := make([]int, C)

	for i := 0; i < A; i++ {
		p[i] = getInt()
	}
	for i := 0; i < B; i++ {
		q[i] = getInt()
	}
	for i := 0; i < C; i++ {
		r[i] = getInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(p)))
	sort.Sort(sort.Reverse(sort.IntSlice(q)))
	sort.Sort(sort.Reverse(sort.IntSlice(r)))

	sumP := 0
	for i := 0; i < X; i++ {
		sumP += p[i]
	}
	sumQ := 0
	for i := 0; i < Y; i++ {
		sumQ += q[i]
	}

	ans := sumQ + sumP
	for i := 0; i < C; i++ {
		ans0 := ans
		if X != 0 {
			ans0 = ans - p[X-1] + r[i]
		}
		ans1 := ans
		if Y != 0 {
			ans1 = ans - q[Y-1] + r[i]
		}
		if ans0 > ans1 {
			if ans0 > ans {
				X--
				ans = ans0
			}
		} else {
			if ans1 > ans {
				Y--
				ans = ans1
			}
		}
	}
	out(ans)
}
