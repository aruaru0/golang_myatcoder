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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	s := getString()
	a := make([]int, N*2)
	if s[0] == 'W' || s[len(s)-1] == 'W' {
		out(0)
		return
	}
	L := 1
	R := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			a[i] = a[i-1]
			if a[i] == 1 {
				R++
			} else {
				L++
			}
		} else {
			if a[i-1] == 1 {
				a[i] = 0
				L++
			} else {
				a[i] = 1
				R++
			}
		}
	}
	// out(a)
	// out(L, R)
	if L != R {
		out(0)
		return
	}
	cnt := 0
	ans := 1
	for i := 0; i < len(s); i++ {
		if a[i] == 0 {
			cnt++
		} else {
			ans *= cnt
			ans %= mod
			cnt--
		}
	}
	for i := 1; i <= N; i++ {
		ans *= i
		ans %= mod
	}
	out(ans)
}
