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

func solve(s []byte) int {
	n := make([]int, 0)
	op := make([]byte, 0)
	x := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			v, _ := strconv.Atoi(string(s[x:i]))
			n = append(n, v)
			op = append(op, s[i])
			x = i + 1
		}
	}
	v, _ := strconv.Atoi(string(s[x:]))
	n = append(n, v)
	// out(n)
	// out(string(op))
	ans := n[0]
	for i := 0; i < len(op); i++ {
		if op[i] == '-' {
			ans -= n[i+1]
		} else {
			ans += n[i+1]
		}
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	s := []byte(getString())

	cnt := 0
	ans := -int(1e9)
	for cnt < len(s) {
		for s[0] < '0' || s[0] > '9' || s[len(s)-1] < '0' || s[len(s)-1] > '9' {
			s = append(s[1:], s[0])
			cnt++
		}
		// out(string(s))
		ans = max(ans, solve(s))
		s = append(s[1:], s[0])
		cnt++
	}

	out(ans)
}
