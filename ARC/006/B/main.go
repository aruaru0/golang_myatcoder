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

var N, L int
var s []string
var y string

func check(p int) bool {
	ret := false
	p = p*2 + 1
	for i := 0; i < L; i++ {
		if s[i][p-1] == '-' {
			p -= 2
		} else if s[i][p+1] == '-' {
			p += 2
		}
	}
	if y[p] == 'o' {
		ret = true
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanLines)
	sc.Buffer([]byte{}, 1000000)
	a := getString()
	fmt.Sscanf(a, "%d %d", &N, &L)
	s = make([]string, L)
	for i := 0; i < L; i++ {
		s[i] = " " + getString() + " "
	}
	y = " " + getString() + " "

	for i := 0; i < N; i++ {
		ret := check(i)
		if ret == true {
			out(i + 1)
			return
		}
	}
}
