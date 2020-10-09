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

func main() {
	sc.Split(bufio.ScanWords)
	n := getInt()
	a := getInts(n)
	m := getInt()
	b := getInts(m)

	u := a[0]
	d := 1
	for i := 1; i < n; i++ {
		d *= a[i]
	}
	for i := 0; i < m; i++ {
		if i%2 == 0 {
			d *= b[i]
		} else {
			u *= b[i]
		}
	}

	if u < 0 && d < 0 {
		u, d = -u, -d
	}
	if u > 0 && d < 0 {
		u, d = -u, -d
	}
	g := gcd(abs(u), abs(d))
	out(u/g, d/g)
}
