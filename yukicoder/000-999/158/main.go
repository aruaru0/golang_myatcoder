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

var d0, d1 int
var b, c []int

func calc(a, b []int, din int) ([]int, int) {
	ret := make([]int, 3)
	for i := a[0]; i >= 0; i-- {
		for j := a[1]; j >= 0; j-- {
			d := din - i*1000 - j*100
			// out(d, din, i, j, a)
			if d >= 0 && d <= a[2] {
				ret[0] = a[0] - i + b[0]
				ret[1] = a[1] - j + b[1]
				ret[2] = a[2] - d + b[2]
				return ret, 1
			}
		}
	}
	return ret, -1
}

var memo [11][101][10001]int

func rec(a []int, s string) int {
	// out("rev", a, s)
	if memo[a[0]][a[1]][a[2]] != -1 {
		return memo[a[0]][a[1]][a[2]]
	}

	n, m := 0, 0

	ret, ok := calc(a, b, d0)
	if ok == 1 {
		n = rec(ret, s+"B") + 1
		// out("---->", n)
	}
	ret, ok = calc(a, c, d1)
	if ok == 1 {
		m = rec(ret, s+"C") + 1
		// out("---->", m)
	}
	memo[a[0]][a[1]][a[2]] = max(n, m)
	return max(n, m)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	a := getInts(3)
	d0 = getInt()
	b = getInts(3)
	d1 = getInt()
	c = getInts(3)

	for i := 0; i < 11; i++ {
		for j := 0; j < 101; j++ {
			for k := 0; k < 10001; k++ {
				memo[i][j][k] = -1
			}

		}
	}

	ret := rec(a, "")
	out(ret)
}
