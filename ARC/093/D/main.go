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

const lim = 100
const lim2 = 50

func main() {
	sc.Split(bufio.ScanWords)
	A, B := getInt(), getInt()
	a := A - 1
	b := B - 1
	// white block
	fmt.Println(((a+lim2-1)/lim2+(b+lim2-1)/lim2)*2+2, lim)
	h := (a + lim2 - 1) / lim2

	for x := 0; x < lim; x++ {
		fmt.Print("#")
	}
	fmt.Println()
	for y := 0; y < h; y++ {
		for x := 0; x < lim; x += 2 {
			if a > 0 {
				fmt.Print(".#")
				a--
			} else {
				fmt.Print("##")
			}
		}
		fmt.Println()
		for x := 0; x < lim; x++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	// Black
	h = (b + lim2 - 1) / lim2
	for y := 0; y < h; y++ {
		for x := 0; x < lim; x++ {
			fmt.Print(".")
		}
		fmt.Println()
		for x := 0; x < lim; x += 2 {
			if b > 0 {
				fmt.Print("#.")
				b--
			} else {
				fmt.Print("..")
			}
		}
		fmt.Println()
	}
	for x := 0; x < lim; x++ {
		fmt.Print(".")
	}
	fmt.Println()
}
