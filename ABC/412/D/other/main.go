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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

type edge struct {
	u, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	f := make(map[edge]bool)

	for i := 0; i < M; i++ {
		u, v := getI()-1, getI()-1
		if u > v {
			u, v = v, u
		}
		f[edge{u, v}] = true
	}
	ans := int(1e10)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}

	for {
		g := make(map[edge]bool)
		for i := 0; i < N; i++ {
			u, v := a[i], a[(i+1)%N]
			if u > v {
				u, v = v, u
			}
			g[edge{u, v}] = true
		}

		cnt0 := 0
		find0 := 0
		for e := range g {
			if !f[e] {
				cnt0++
			} else {
				find0++
			}
		}
		cnt0 += len(f) - find0
		// for i := 0; i < N; i++ {
		// 	for j := 0; j < N; j++ {
		// 		if f[edge{i, j}] != g[edge{i, j}] {
		// 			cnt0++
		// 		}
		// 	}
		// }

		ans = min(ans, cnt0)

		for d := 3; d <= N-3; d++ {
			h := make(map[edge]bool)
			for i := 0; i < d; i++ {
				u, v := a[i], a[(i+1)%d]
				if u > v {
					u, v = v, u
				}
				h[edge{u, v}] = true
			}
			for i := 0; i < N-d; i++ {
				u, v := a[i+d], a[(i+1)%(N-d)+d]
				if u > v {
					u, v = v, u
				}
				h[edge{u, v}] = true
			}
			cnt1 := 0
			find1 := 0
			for e := range h {
				if !f[e] {
					cnt1++
				} else {
					find1++
				}
			}
			cnt1 += len(f) - find1
			ans = min(ans, cnt1)
		}

		if NextPermutation(sort.IntSlice(a)) == false {
			break
		}
	}

	out(ans)
}
