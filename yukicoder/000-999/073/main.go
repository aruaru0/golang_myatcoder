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

func nCr(n, k int) int {
	ret := 1
	for i := 0; i < k; i++ {
		ret *= n - i
	}
	for i := 1; i <= k; i++ {
		ret /= i
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	a := getInts(26)
	h := a[int('h'-'a')]
	e := a[int('e'-'a')]
	l := a[int('l'-'a')]
	o := a[int('o'-'a')]
	w := a[int('w'-'a')]
	r := a[int('r'-'a')]
	d := a[int('d'-'a')]
	// out(h, e, l, o, w, r, d)
	ans := h * e * w * r * d
	o -= 2
	sum := 0
	for i := 0; i <= o; i++ {
		x := nCr(i+1, 1) * nCr(o-i+1, 1)
		sum = max(sum, x)
	}
	ans *= sum
	sum = 0
	l -= 3
	for i := 0; i <= l; i++ {
		x := nCr(i+2, 2) * nCr(l-i+1, 1)
		sum = max(sum, x)
	}
	ans *= sum
	out(ans)
}
