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

func f(s string) int {
	N, _ := strconv.Atoi(s)
	a := 0
	for i := 0; i <= N; i++ {
		sum := 0
		n := i
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		a = max(a, sum)
	}
	return a
}

func solve(s string) int {
	// １桁
	if len(s) == 1 {
		v, _ := strconv.Atoi(s)
		return v
	}
	// ２桁目以降がすべて９
	flg := true
	for i := 1; i < len(s); i++ {
		if s[i] != '9' {
			flg = false
		}
	}
	x := (len(s)-1)*9 + int(s[0]-'0') - 1
	if flg == true {
		x++
	}

	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := getString()

	out(solve(s))

	// for i := 1; i < 10000; i++ {
	// 	s := fmt.Sprintf("%d", i)
	// 	x := solve(s)
	// 	y := f(s)
	// 	if x != y {
	// 		out(s, x, y, "------")
	// 	} else {
	// 		out(s, x, y)
	// 	}
	// }
}
