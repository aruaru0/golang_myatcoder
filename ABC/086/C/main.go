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

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	t := make([]int, N)
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		t[i], x[i], y[i] = getInt(), getInt(), getInt()
	}

	cx := 0
	cy := 0
	ct := 0
	ok := true
	for i := 0; i < N; i++ {
		dx := abs(cx - x[i])
		dy := abs(cy - y[i])
		dt := t[i] - ct
		l := dx + dy
		if dt < l {
			ok = false
			break
		}
		l -= dt
		if l%2 != 0 {
			ok = false
		}
		cx = x[i]
		cy = y[i]
		ct = t[i]
		// out(dx, dy, dt)
	}
	if ok {
		out("Yes")
	} else {
		out("No")
	}

}
