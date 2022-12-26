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

func main() {
	sc.Split(bufio.ScanWords)
	N, M, V, P := getInt(), getInt(), getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	th := a[P-1]
	l := 0
	r := N
	for l+1 != r {
		m := (l + r) / 2
		val := a[m] + M
		cnt := M * (V - 1)
		if val < th {
			r = m
			continue
		}
		for i, v := range a {
			if i == m {
				continue
			}
			if i < P && v != th {
				cnt -= M
			} else {
				cnt -= min(M, val-v)
			}
		}
		if cnt <= 0 {
			l = m
		} else {
			r = m
		}
	}
	out(l + 1)
}
