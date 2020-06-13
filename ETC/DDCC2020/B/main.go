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
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	a := make([]int, N)
	sum := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		sum += a[i]
	}

	l := 0
	r 
	:= 0
	m := 0
	c := 0
	for i := 0; i < N; i++ {
		if r < sum/2 {
			l = r
			r += a[i]
			c++
		} else {
			m = a[c-1]
			r = sum - r
			break
		}
	}
	if N == 2 {
		l = a[0]
		r = a[1]
		m = 0
	}

	//	out(l, r, m, c)
	ans := sum * 2
	if l+m >= r {
		ans = min(ans, (l + m - r))
	} else {
		ans = min(ans, r-(l+m))
	}
	if r+m >= l {
		ans = min(ans, (r + m - l))
	} else {
		ans = min(ans, (l - (r + m)))
	}
	out(ans)

}
