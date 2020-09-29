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

	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
	}

	// for i := 1; i <= N; i++ {
	// 	for j := 1; j <= N; j++ {
	// 		a[i-1][j-1] = (2*i - j) % N
	// 		if a[i-1][j-1] <= 0 {
	// 			a[i-1][j-1] += N
	// 		}
	// 	}
	// }
	for i := 0; i < N; i++ {
		x := i
		y := i

		for j := 0; j < N; j++ {
			a[y][x] = i + 1
			x--
			y++
			if x < 0 {
				x += N
			}
			y %= N
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fprint(w, a[i][j], " ")
		}
		fmt.Fprintln(w)
	}
}
