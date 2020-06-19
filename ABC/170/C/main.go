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
	X, N := getInt(), getInt()
	p := make([]int, 102)
	for i := 0; i < N; i++ {
		pos := getInt()
		p[pos] = 1
	}
	if p[X] == 0 {
		out(X)
		return
	}
	p0 := 101
	for i := X; i <= 101; i++ {
		if p[i] == 0 {
			p0 = i
			break
		}
	}
	p1 := 0
	for i := X; i >= 0; i-- {
		if p[i] == 0 {
			p1 = i
			break
		}
	}
	if abs(p0-X) < abs(p1-X) {
		out(p0)
	} else {
		out(p1)
	}
}
