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
	x := getString()
	c := make([]int, 10)
	for i := 0; i < len(x); i++ {
		p := int(x[i] - '0')
		c[p]++
	}

	for i := 0; i < 10; i++ {
		if c[i] == len(x) {
			out(-1)
			return
		}
	}
	if c[0] == len(x)-1 {
		out(-1)
		return
	}

	s := make([]byte, 0, len(x))
	for i := 9; i >= 0; i-- {
		for j := 0; j < c[i]; j++ {
			s = append(s, byte(i+'0'))
		}
	}

	pos := 0
	for i := len(s) - 1; i >= 1; i-- {
		if s[i] != s[i-1] {
			pos = i
			break
		}
	}

	s[pos], s[pos-1] = s[pos-1], s[pos]
	out(string(s))

}
