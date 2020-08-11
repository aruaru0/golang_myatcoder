package main

import (
	"bufio"
	"fmt"
	"math"
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

func lowerBoundP(a []pair, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x >= x
	})
	return idx
}

func upperBoundP(a []pair, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].x > x
	})
	return idx
}

type pair struct {
	x, i int
}

func calc(a, b, x, r, h int) float64 {
	s := float64(r) * float64(r) * math.Pi * float64(h) / 3
	if a > x {
		H := float64(h - (a - x))
		R := float64(r) * H / float64(h)
		s = R * R * math.Pi * H / 3
	}
	if b < x+h {
		H := float64(h - (b - x))
		R := float64(r) * H / float64(h)
		s -= R * R * math.Pi * H / 3
	}
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	N, Q := getInt(), getInt()
	x := make([]int, N)
	r := make([]int, N)
	h := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], r[i], h[i] = getInt(), getInt(), getInt()
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < Q; i++ {
		A, B := getInt(), getInt()
		ans := 0.0
		for j := 0; j < N; j++ {
			if B <= x[j] || A >= x[j]+h[j] {
				continue
			}
			ans += calc(A, B, x[j], r[j], h[j])
		}
		fmt.Fprintln(w, ans)
	}
}
