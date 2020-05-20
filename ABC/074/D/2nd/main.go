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
	a := make([][]int, N)
	b := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
		b[i] = make([]int, N)
		for j := 0; j < N; j++ {
			a[i][j] = getInt()
			b[i][j] = a[i][j]
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if b[i][j] > b[i][k]+b[k][j] {
					out(-1)
					return
				}
			}
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if a[i][j] == a[i][k]+a[k][j] {
					if i != k && j != k {
						b[i][j] = 0
					}
				}
			}
		}
	}
	// out("----")
	// for i := 0; i < N; i++ {
	// 	out(b[i])
	// }

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans += b[i][j]
		}
	}
	out(ans / 2)
}
