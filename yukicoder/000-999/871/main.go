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

type flog struct {
	x, a int
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()-1
	x := getInts(N)
	a := getInts(N)
	// x = append([]int{0}, x...)
	// a = append([]int{0}, a...)

	lpos, rpos := K, K
	lval, rval := x[K]-a[K], x[K]+a[K]

	flg := true
	for flg {
		flg = false
		for lpos-1 >= 0 && lval <= x[lpos-1] {
			flg = true
			lpos--
			lval = min(lval, x[lpos]-a[lpos])
			rval = max(rval, x[lpos]+a[lpos])
		}
		for rpos+1 < N && x[rpos+1] <= rval {
			flg = true
			rpos++
			lval = min(lval, x[rpos]-a[rpos])
			rval = max(rval, x[rpos]+a[rpos])
		}
	}
	out(rpos - lpos + 1)
}
