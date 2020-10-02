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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	a := getInts(N)

	if a[0] != a[N-1] {
		m := make(map[int]int)
		col := a[0]
		m[col]++
		for i := 1; i < N; i++ {
			if a[i] != col {
				col = a[i]
				if m[col] != 0 {
					out(-1)
					return
				}
				m[col]++
			}
		}
		out(0)
		return
	} else {
		l := 0
		for a[l] == a[0] {
			l++
			if l == N {
				out(0)
				return
			}
		}
		r := N - 1
		for a[r] == a[0] {
			r--
		}
		m := make(map[int]int)
		col := a[0]
		m[col]++
		for i := l; i <= r; i++ {
			if a[i] != col {
				col = a[i]
				if m[col] != 0 {
					out(-1)
					return
				}
				m[col]++
			}
		}
		out(1)
	}
}
