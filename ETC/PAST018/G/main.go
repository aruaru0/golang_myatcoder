package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

type pair struct {
	a, b int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	a := getInts(n)
	orgb := getInts(n)

	b := make([]int, n)
	copy(b, orgb)
	if f(n, a, b) {
		fmt.Println("Yes")
		return
	}
	copy(b, orgb)
	if f(n, b, a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func f(n int, a, b []int) bool {
	cnt := 0
	for i := 0; i < n-1; i++ {
		if a[i] == b[i] {
			continue
		} else {
			b[i], b[i+1] = b[i+1], b[i]
			cnt++
		}
	}

	// 変更点が２箇所の場合OKする
	if cnt == 0 || cnt == 2 {
		ans := true
		for i := range a {
			if a[i] != b[i] {
				ans = false
				break
			}
		}
		return ans
	} else if cnt == 1 { // 変更が１箇所の場合
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		//　同じものが存在すればOK
		for i := 0; i < n-1; i++ {
			if a[i] == a[i+1] {
				return true
			}
		}
		return false
	}
	return false
}
