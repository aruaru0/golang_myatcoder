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

func conv(s string) []string {
	ret := make([]string, 0)
	p := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '*' || s[i] == '+' {
			ret = append(ret, s[p:i])
			ret = append(ret, string(s[i]))
			p = i + 1
		}
	}
	ret = append(ret, s[p:])
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := getString()
	r := conv(s)
	// out(r)
	ans, _ := strconv.Atoi(r[0])
	for i := 1; i < len(r); i += 2 {
		x, _ := strconv.Atoi(r[i+1])
		if r[i] == "*" {
			ans += x
		} else {
			ans *= x
		}
	}
	out(ans)
}
