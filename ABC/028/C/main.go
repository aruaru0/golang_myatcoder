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
	a := make([]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = getInt()
	}

	n := 1 << 5
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		x := 0
		cnt := 0
		for j := 0; j < 5; j++ {
			if i>>uint(j)%2 == 1 {
				x += a[j]
				cnt++
			}
		}
		if cnt == 3 {
			b = append(b, x)
		}
	}
	sort.Ints(b)
	out(b[len(b)-3])
}
