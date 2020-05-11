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
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getString()
	}
	c := make([][2]int, N)
	for i := 0; i < N; i++ {
		for _, v := range s[i] {
			if v == '(' {
				c[i][0]++
			} else {
				if c[i][0] == 0 {
					c[i][1]++
				} else {
					c[i][0]--
				}
			}
		}
	}
	sort.Slice(c, func(i, j int) bool {
		if c[i][1] == 0 {
			return true
		}
		if c[j][1] == 0 {
			return false
		}
		if c[i][0] == 0 {
			return false
		}
		if c[j][0] == 0 {
			return true
		}
		di := c[i][0] - c[i][1]
		dj := c[j][0] - c[j][1]
		return di > dj
	})

	if c[0][1] != 0 {
		out("No")
		return
	}

	// out(c)

	l := c[0][0]
	r := c[0][1]
	for i := 1; i < N; i++ {
		x := min(l, c[i][1])
		y := min(r, c[i][0])
		l = l - x + c[i][0] - y
		r = r - y + c[i][1] - x
		if r > 0 {
			out("No")
			return
		}
		// out(c[i], x, y, l, r)
	}
	// if l == 0 && r == 0 && cntl == cntr {
	if l == 0 && r == 0 {
		out("Yes")
	} else {
		out("No")
	}
}
