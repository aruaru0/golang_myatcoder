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

var m map[int]bool
var N, K int

func check(a []int) []int {
	n := len(a)
	pat := [][2]int{{1, 6}, {2, 5}, {3, 4}}
	for i := n - 1; i >= 0; i-- {
		for k := 0; k < 3; k++ {
			x := i + pat[k][0]
			y := i + pat[k][1]
			if x >= n || y >= n {
				continue
			}
			if a[x] == 1 && a[y] == 1 {
				a[i] = 1
			}
		}
	}
	return a
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K = getI(), getI()
	a := getInts(K)

	m = make(map[int]bool)
	for i := 0; i < K; i++ {
		m[a[i]] = true
	}

	ng := false
	if a[K-1] >= 100 {
		for i := 0; i < K; i++ {
			for p := 1; p <= 6; p++ {
				if a[i] < 50 {
					continue
				}
				start := a[i] - p
				if m[start] {
					continue
				}
				if start <= 0 {
					continue
				}
				pos0 := a[i]
				pos1 := start + 7 - p
				if m[pos0] && m[pos1] {
					ng = true
				}
				// out(a[i], p, "start", start, "to", pos0, pos1)
			}
		}
	}
	if ng {
		out("No")
		return
	}

	x := make([]int, 100)
	for i := 0; i < K; i++ {
		if a[i] < 100 {
			x[a[i]] = 1
		}
	}
	check(x)
	if x[1] == 0 {
		out("Yes")
	} else {
		out("No")
	}
}
