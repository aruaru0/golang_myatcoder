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

func check(a, b string) (byte, bool) {
	var m [26]int
	for _, v := range a {
		m[int(v-'a')]++
	}
	for _, v := range b {
		if m[int(v-'a')] != 0 {
			return byte(v), true
		}
	}
	return ' ', false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	a := make([]string, N)
	for i := 0; i < N; i++ {
		a[i] = getString()
	}

	s := ""
	for i := 0; i < N; i++ {
		// out(a[i], a[N-1-i])
		c, flg := check(a[i], a[N-1-i])
		if !flg {
			out("-1")
			return
		}
		s = s + string(c)
	}
	out(s)
}
