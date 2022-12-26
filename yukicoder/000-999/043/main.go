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

type pair struct {
	win, x int
}

type card struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getString()
	}
	x := make([]pair, N)
	y := make([]card, 0)
	for i := 0; i < N; i++ {
		x[i].x = i
		for j := i; j < N; j++ {
			if s[i][j] == '#' {
				continue
			}
			if s[i][j] == 'o' {
				x[i].win++
			}
			if s[i][j] == 'x' {
				x[j].win++
			}
			if s[i][j] == '-' {
				y = append(y, card{i, j})
			}
		}
	}

	n := 1 << uint(len(y))
	// out("n", n)
	ans := N
	for i := 0; i < n; i++ {
		xd := make([]pair, N)
		copy(xd, x)
		for j := 0; j < len(y); j++ {
			if (i>>j)%2 == 1 {
				xd[y[j].x].win++
			} else {
				xd[y[j].y].win++
			}
		}
		// fmt.Printf("%b\n", i)
		// out(xd)

		sort.Slice(xd, func(i, j int) bool {
			if xd[i].win == xd[j].win {
				return xd[i].x < xd[j].x
			}
			return xd[i].win > xd[j].win
		})
		// out(xd)
		no := 0
		prev := -1
		for i := 0; i < N; i++ {
			if prev != xd[i].win {
				prev = xd[i].win
				no++
			}
			if xd[i].x == 0 {
				ans = min(ans, no)
			}
		}
	}

	out(ans)
}
