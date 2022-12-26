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

func myAtoi(s string, b []byte) int {
	ret := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < 10; j++ {
			if s[i] == b[j] {
				ret *= 10
				ret += j
				break
			}
		}
	}
	return ret
}

func myItoa(n int, b []byte) string {
	var ret []byte
	for n > 0 {
		x := n % 10
		ret = append(ret, b[x])
		n /= 10
	}
	N := len(ret)
	for i := 0; i < N/2; i++ {
		ret[i], ret[N-1-i] = ret[N-1-i], ret[i]
	}
	return string(ret)
}

func main() {
	sc.Split(bufio.ScanWords)

	b := make([]byte, 10)
	for i := 0; i < 10; i++ {
		b[i] = []byte(getString())[0]
	}
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		t := getString()
		a[i] = myAtoi(t, b)
	}
	sort.Ints(a)
	for i := 0; i < N; i++ {
		out(myItoa(a[i], b))
	}
}
