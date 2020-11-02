package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
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

func solve(h int, w []int) {
	ans := abs(w[0] - h)
	for i := 0; i < len(w); i++ {
		ans = min(ans, abs(h-w[i]))
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	h := getInts(N)
	w := getInts(M)

	if N == 1 {
		solve(h[0], w)
		return
	}
	sort.Ints(h)

	L := make([]int, N+5)
	R := make([]int, N+5)
	sumL := 0
	sumR := 0
	for i := 0; i < N-1; i++ {
		if i%2 == 0 {
			sumL += abs(h[i] - h[i+1])
		}
		L[i] = sumL
	}
	for i := 1; i < N; i++ {
		if i%2 == 1 {
			sumR += abs(h[i] - h[i+1])
		}
		R[i] = sumR
	}

	off := 5
	L = append([]int{0, 0, 0, 0, 0}, L...)
	R = append([]int{0, 0, 0, 0, 0}, R...)

	ans := int(1e18)
	for i := 0; i < M; i++ {
		pos := lowerBound(h, w[i])
		if pos%2 == 0 {
			tot := abs(w[i]-h[pos]) + L[pos-1+off] + (R[N-1+off] - R[pos+off])
			ans = min(ans, tot)
		} else {
			tot := abs(w[i]-h[pos-1]) + L[pos-2+off] + (R[N-1+off] - R[pos-1+off])
			ans = min(ans, tot)
		}
	}
	out(ans)

}
