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

func makeEven(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	x, y := n/2-1, 1
	cnt := 0
	for cnt != n*n {
		prex := x
		prey := y
		x++
		y--
		if x >= n {
			x = 0
		}
		if y < 0 {
			y = n - 1
		}
		if a[y][x] != 0 {
			x = prex
			y = prey + 1
			if y >= n {
				y = 0
			}
		}
		cnt++
		a[y][x] = cnt
	}
	return a
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

func fill(b [][]int, x, y, t int, n []int) {
	x = x * 2
	y = y * 2
	b[y][x] = n[0] + t
	b[y][x+1] = n[1] + t
	b[y+1][x] = n[2] + t
	b[y+1][x+1] = n[3] + t
}

func make4n2(n int) [][]int {
	nn := (n - 2) / 4
	nnn := 2*nn + 1
	a := makeEven(nnn)
	for y := 0; y < nnn; y++ {
		for x := 0; x < nnn; x++ {
			a[y][x]--
			a[y][x] *= 4
		}
	}

	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = make([]int, n)
	}

	for y := 0; y < nnn; y++ {
		for x := 0; x < nnn; x++ {
			if y < nn {
				fill(b, x, y, a[y][x], []int{4, 1, 2, 3})
			} else if y == nn {
				if x == nn {
					fill(b, x, y, a[y][x], []int{1, 4, 2, 3})
				} else {
					fill(b, x, y, a[y][x], []int{4, 1, 2, 3})
				}
			} else if y == nn+1 {
				if x == nn {
					fill(b, x, y, a[y][x], []int{4, 1, 2, 3})
				} else {
					fill(b, x, y, a[y][x], []int{1, 4, 2, 3})
				}
			} else {
				fill(b, x, y, a[y][x], []int{1, 4, 3, 2})
			}
		}
	}

	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()

	var a [][]int

	if N%2 == 1 {
		a = makeEven(N)
	}
	if N%4 == 0 {
		a = make44(N)
	}
	if N%4 == 2 {
		a = make4n2(N)
	}

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			fmt.Print(a[y][x], " ")
		}
		out()
	}
}
