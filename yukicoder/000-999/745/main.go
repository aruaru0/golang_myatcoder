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

func f(a, s int) int {
	ret := 0
	if a > 300 {
		x := a - 300
		ret += s * 2 * 2 * 2 * x
		a -= x
		out(x, a, ret)
	}
	if a > 200 {
		x := a - 200
		ret += s * 2 * 2 * x
		a -= x
		out(x, a, ret)
	}
	if a > 100 {
		x := a - 100
		ret += s * 2 * x
		a -= x
		out(x, a, ret)
	}
	out(a)
	ret += s * a
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, _, d := getInt(), getInt(), getInt(), getInt()

	if d == 10 {
		out("Impossible")
		return
	}
	x := make([]int, a+b)
	for i := 0; i < b; i++ {
		x[i] = 50
	}
	for i := 0; i < a; i++ {
		x[b+i] = 100
	}

	score := 0
	rate := 1
	for i := 0; i < len(x); i++ {
		if i != 0 && i%100 == 0 {
			rate *= 2
		}
		score += x[i] * rate
	}
	out("Possible")
	out(score)
}
