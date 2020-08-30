package main

import (
	"bufio"
	"fmt"
	"math"
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
	N := getInt()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getString()
	}
	code := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ans := math.MaxInt64
	for i := 0; i < N; i++ {
		c := '0'
		for _, e := range s[i] {
			if e > c {
				c = e
			}
		}
		n := 2
		for j := 1; j < len(code); j++ {
			if code[j] == byte(c) {
				n = j + 1
				break
			}
		}
		x, _ := strconv.ParseInt(s[i], n, 64)
		// out(s[i], n, x)
		ans = min(ans, int(x))
	}
	out(ans)
}
