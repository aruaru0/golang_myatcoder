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

func solve(L, R int) int {
	if L == 0 {
		return 0
	}
	l := (L - 1) / 2019
	r := R / 2019

	if r-l >= 1 {
		return 0
	}

	ret := 2019
	for i := L; i < R; i++ {
		for j := i + 1; j <= R; j++ {
			ret = min(ret, (i*j)%2019)
		}
	}
	return ret
}

func ref(L, R int) (int, int, int) {
	ret := 2019
	l := L
	r := R
	for i := L; i < R; i++ {
		for j := i + 1; j <= R; j++ {
			if i*j%2019 < ret {
				l = i
				r = j
				ret = min(ret, (i*j)%2019)
			}
		}
	}
	return ret, l, r
}

// Pfs :　素因数分解し、スライスを作成
func Pfs(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
func main() {
	sc.Split(bufio.ScanWords)
	L, R := getInt(), getInt()
	out(solve(L, R))
}
