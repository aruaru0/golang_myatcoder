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
	grundy := make([]int, 100+1)
	grundy[1] = 0
	for i := 2; i <= 100; i++ {
		var s [4]int
		var x, y int
		x = grundy[i/2] ^ grundy[(i+1)/2]
		s[x]++
		if i >= 3 {
			y = grundy[i/3] ^ grundy[(i+1)/3] ^ grundy[(i+2)/3]
			s[y]++
		}
		c := 0
		for i := 0; i < 4; i++ {
			if s[i] != 0 {
				c++
			} else {
				break
			}
		}
		grundy[i] = c
	}

	if grundy[N] == 0 {
		out("B")
		return
	}
	out("A")
}
