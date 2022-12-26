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

var c []int
var a []int

func dfs(cur int) bool {
	x := c[cur]
	if c[x] != -1 {
		return c[x] == a[cur]
	}
	c[x] = a[cur]
	ret := dfs(x)
	c[x] = -1
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a = make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt() - 1
	}
	c = make([]int, N)
	used := make([]bool, N)
	for i := 0; i < N; i++ {
		if a[i] == i {
			c[i] = i
			used[i] = true
		} else {
			c[i] = -1
		}
	}

	for i := 0; i < N; i++ {
		if c[i] == -1 {
			for j := 0; j < N; j++ {
				if used[j] == false {
					if c[j] != -1 && c[j] != a[i] {
						continue
					}
					t := c[j]
					c[i] = j
					c[j] = a[i]
					// out(i, j, dfs(j))
					if dfs(j) {
						break
					}
					c[i] = -1
					c[j] = t
				}
			}
			if c[i] == -1 {
				out("No")
				return
			}
		}
	}
	out("Yes")
}
