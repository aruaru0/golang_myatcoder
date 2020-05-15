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
	var c [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = getInt()
		}
	}

	var a [3]int
	var b [3]int
	for a[0] = 0; a[0] <= 100; a[0]++ {
		b[0] = c[0][0] - a[0]
		b[1] = c[0][1] - a[0]
		b[2] = c[0][2] - a[0]
		for a[1] = 0; a[1] <= 100; a[1]++ {
			for a[2] = 0; a[2] <= 100; a[2]++ {
				ans := true
			L0:
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {
						if c[i][j] != a[i]+b[j] {
							ans = false
							break L0
						}
					}
				}
				if ans == true {
					out("Yes")
					return
				}
			}
		}
	}
	out("No")
}
