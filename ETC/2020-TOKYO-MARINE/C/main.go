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

func f(N int, a []int) ([]int, bool) {
	b := make([]int, N+1)
	x := N
	flg := false
	for i := 0; i < N; i++ {
		if a[i] != x {
			flg = true
		}
		f := max(0, i-a[i])
		t := min(N, i+a[i]+1)
		b[f]++
		b[t]--
	}
	for i := 1; i <= N; i++ {
		b[i] += b[i-1]
	}
	return b, flg
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	a := make([]int, N+1)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	flg := true
	for i := 0; i < K; i++ {
		a, flg = f(N, a)
		if flg == false {
			// out(i)
			break
		}
	}
	ans := a
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < N; i++ {
		fmt.Fprint(w, ans[i], " ")
	}
	fmt.Fprintln(w)
}
