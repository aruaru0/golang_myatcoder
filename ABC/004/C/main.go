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
	N := getInt()

	b := make([][]int, 0)
	a := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i <= N; i++ {
		c := make([]int, 6)
		copy(c, a)
		b = append(b, c)
		x, y := i%5, i%5+1
		a[x], a[y] = a[y], a[x]
		flg := true
		for j := 0; j < 6; j++ {
			if a[j] != j+1 {
				flg = false
				break
			}
		}
		if flg {
			break
		}
	}
	x := N % 30
	for i := 0; i < 6; i++ {
		fmt.Print(b[x][i])
	}
	out()
}
