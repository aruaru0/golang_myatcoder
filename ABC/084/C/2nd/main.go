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
	c := make([]int, N-1)
	s := make([]int, N-1)
	f := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		c[i], s[i], f[i] = getInt(), getInt(), getInt()
	}
	// out(c, s, f)
	for i := 0; i < N-1; i++ {
		t := 0
		for j := i; j < N-1; j++ {
			// out(":", t)
			if t < s[j] {
				t = (s[j]) + c[j]
			} else if t%f[j] == 0 {
				t += c[j]
			} else {
				t += (f[j] - t%f[j]) + c[j]
			}
		}
		out(t)
	}
	out(0)

}
