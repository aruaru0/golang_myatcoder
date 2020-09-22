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

func make44(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	cnt := 1
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			switch y % 4 {
			case 0:
				fallthrough
			case 3:
				if x%4 == 0 || x%4 == 3 {
					a[y][x] = cnt
				}

			case 1:
				fallthrough
			case 2:
				if x%4 == 1 || x%4 == 2 {
					a[y][x] = cnt
				}
			}
			cnt++
		}
	}

	cnt = n * n
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			switch y % 4 {
			case 1:
				fallthrough
			case 2:
				if x%4 == 0 || x%4 == 3 {
					a[y][x] = cnt
				}

			case 0:
				fallthrough
			case 3:
				if x%4 == 1 || x%4 == 2 {
					a[y][x] = cnt
				}
			}
			cnt--
		}
	}

	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	n, x, y, z := getInt(), getInt()-1, getInt()-1, getInt()
	a := make44(n)

	b := (a[y][x] - 1) ^ (z - 1)

	// for y := 0; y < n; y++ {
	// 	for x := 0; x < n; x++ {
	// 		fmt.Print(a[y][x], " ")
	// 	}
	// 	out()
	// }
	// out("---")
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			fmt.Print(((a[y][x]-1)^b)+1, " ")
		}
		out()
	}
}
