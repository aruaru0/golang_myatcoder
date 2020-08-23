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

type pair struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, 0)
	for i := 0; i < N; i++ {
		x := getInt()
		a = append(a, x)
	}
	s := make([]pair, 0)
	sort.Ints(a)
	mi := a[0]
	ma := a[N-1]
	a = a[1 : N-1]
	for _, e := range a {
		if e >= 0 {
			s = append(s, pair{mi, e})
			mi -= e
		} else {
			s = append(s, pair{ma, e})
			ma -= e
		}
	}

	s = append(s, pair{ma, mi})
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, ma-mi)
	for _, e := range s {
		fmt.Fprintln(w, e.x, e.y)
	}
}
