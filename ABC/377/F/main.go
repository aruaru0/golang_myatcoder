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

type P struct {
	f, s int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	X := make(map[int]bool)
	Y := make(map[int]bool)
	S := make(map[int]bool)
	D := make(map[int]bool)
	for i := 0; i < m; i++ {
		a, b := getI()-1, getI()-1
		X[a] = true
		Y[b] = true
		S[a+b] = true
		D[a-b] = true
	}

	ans := n * n
	ans -= len(X) * n
	ans -= len(Y) * n
	for s := range S {
		ans -= n - abs(s-(n-1))
	}
	for d := range D {
		ans -= n - abs(d)
	}

	cnt := make(map[P]int)
	add := func(x, y int) {
		if x < 0 || y < 0 || x >= n || y >= n {
			return
		}
		cnt[P{x, y}]++
	}

	for x := range X {
		for y := range Y {
			add(x, y)
		}
	}
	for x := range X {
		for s := range S {
			add(x, s-x)
		}
	}

	for x := range X {
		for d := range D {
			add(x, x-d)
		}
	}

	for y := range Y {
		for s := range S {
			add(s-y, y)
		}
	}
	for y := range Y {
		for d := range D {
			add(y+d, y)
		}
	}

	for s := range S {
		for d := range D {
			if (s+d)%2 != 0 {
				continue
			}
			add((s+d)/2, (s-d)/2)
		}
	}

	for p := range cnt {
		num := cnt[p]
		if num == 1 {
			ans += 1
		}
		if num == 3 {
			ans += 2
		}
		if num == 6 {
			ans += 3
		}
	}
	out(ans)

}
