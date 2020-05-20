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

func check(c []int, s string) bool {
	N := len(s)
	for i := 1; i < N; i++ {
		if c[i] == 0 {
			if s[i] == 'o' {
				c[i+1] = c[i-1]
			} else {
				c[i+1] = c[i-1] ^ 1
			}
		} else {
			if s[i] == 'x' {
				c[i+1] = c[i-1]
			} else {
				c[i+1] = c[i-1] ^ 1
			}
		}
	}

	if c[0] != c[N] {
		return false
	}
	if c[0] == 0 {
		if s[0] == 'o' {
			if c[1] != c[N-1] {
				return false
			}
		} else {
			if c[1] == c[N-1] {
				return false
			}
		}
	}
	if c[0] == 1 {
		if s[0] == 'o' {
			if c[1] == c[N-1] {
				return false
			}
		} else {
			if c[1] != c[N-1] {
				return false
			}
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < N; i++ {
		if c[i] == 0 {
			fmt.Fprint(w, "S")
		} else {
			fmt.Fprint(w, "W")
		}
	}
	fmt.Fprintln(w)

	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	s := getString()

	c := make([]int, N+1)

	c[0] = 0
	c[1] = 0
	ret := check(c, s)
	// out(ret, c)
	// out(s)
	if ret {
		return
	}

	c[0] = 0
	c[1] = 1
	ret = check(c, s)
	// out(ret, c)
	if ret {
		return
	}

	c[0] = 1
	c[1] = 0
	ret = check(c, s)
	// out(ret, c)
	if ret {
		return
	}

	c[0] = 1
	c[1] = 1
	ret = check(c, s)
	// out(ret, c)
	if ret {
		return
	}

	out(-1)
}
