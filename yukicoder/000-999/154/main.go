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

func solve(s string) {
	w, g, r := 0, 0, 0
	wst := 0 // last W->G->R state
	for _, e := range s {
		if e == 'W' {
			w++
			wst = 2
		}
		if e == 'G' {
			g++
			if wst == 2 {
				wst = 1
			}
		}
		if e == 'R' {
			r++
			if wst == 1 {
				wst = 0
			}
		}
		if w >= g && g >= r {
			continue
		}
		out("impossible")
		return
	}

	if wst == 0 && w >= g && g == r {
		out("possible")
		return
	}
	out("impossible")
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	T := getInt()
	for i := 0; i < T; i++ {
		s := getString()
		solve(s)
	}
}
