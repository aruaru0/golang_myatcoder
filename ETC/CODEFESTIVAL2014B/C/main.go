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
	s1 := getString()
	s2 := getString()
	s3 := getString()
	m := make([][26]int, 3)
	for _, e := range s1 {
		x := int(e - 'A')
		m[0][x]++
	}
	for _, e := range s2 {
		x := int(e - 'A')
		m[1][x]++
	}
	for _, e := range s3 {
		x := int(e - 'A')
		m[2][x]++
	}
	n := len(s1)
	ma, mi := 0, 0
	for i := 0; i < 26; i++ {
		ma += max(0, m[2][i]-m[1][i])
		mi += min(m[0][i], m[2][i])
	}

	if ma < n/2 && n/2 <= mi {
		out("YES")
		return
	}
	out("NO")
}
