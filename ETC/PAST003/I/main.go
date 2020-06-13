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
	N := getInt()
	Q := getInt()

	r := make([]int, N)
	c := make([]int, N)
	for j := 0; j < N; j++ {
		r[j] = j
		c[j] = j
	}

	rot := false
	for i := 0; i < Q; i++ {
		op := getInt()
		if op == 1 {
			a, b := getInt()-1, getInt()-1
			if rot {
				c[a], c[b] = c[b], c[a]
			} else {
				r[a], r[b] = r[b], r[a]
			}
			// out("op1", a, b)
		} else if op == 2 {
			a, b := getInt()-1, getInt()-1
			if rot {
				r[a], r[b] = r[b], r[a]
			} else {
				c[a], c[b] = c[b], c[a]
			}
			// out("op2", a, b)
		} else if op == 3 {
			rot = !rot
			// out("op3")
		} else {
			a, b := getInt()-1, getInt()-1
			if rot {
				b, a = a, b
			}
			out(N*r[a] + c[b])
		}

		// out("----", rot, r, c)
		// for j := 0; j < N; j++ {
		// 	for k := 0; k < N; k++ {
		// 		a, b := j, k
		// 		if rot {
		// 			a, b = k, j
		// 		}
		// 		fmt.Print(N*r[a] + c[b])
		// 	}
		// 	out()
		// }
	}
}
