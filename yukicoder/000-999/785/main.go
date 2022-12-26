package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func f(a string) int {
	if a == "NONE" {
		return 256
	}

	x := strings.Split(a, ",")
	b := make([]bool, 16)
	for _, e := range x {
		v, _ := strconv.ParseInt(string(e), 16, 64)
		b[v] = true
	}
	cnt := 0
	for i := 0; i < 256; i++ {
		if b[i%16] == false && b[i/16] == false {
			cnt++
		}
	}
	return cnt
}

func main() {
	sc.Split(bufio.ScanLines)
	r := getString()
	g := getString()
	b := getString()

	c := f(r)
	c *= f(g)
	c *= f(b)
	out(c)
}
