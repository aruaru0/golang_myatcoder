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

var N int

var memo [110][110][110]float64

func rec(n3, n2, n1 int) float64 {
	if memo[n3][n2][n1] >= 0 {
		return memo[n3][n2][n1]
	}

	n := float64(N) / float64(n1+n2+n3)
	if n1 > 0 {
		n += rec(n3, n2, n1-1) * float64(n1) / float64(n1+n2+n3)
	}
	if n2 > 0 {
		n += rec(n3, n2-1, n1+1) * float64(n2) / float64(n1+n2+n3)
	}
	if n3 > 0 {
		n += rec(n3-1, n2+1, n1) * float64(n3) / float64(n1+n2+n3)
	}
	memo[n3][n2][n1] = n
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	a := getInts(N)
	n3, n2, n1 := 0, 0, 0
	for _, v := range a {
		switch v {
		case 0:
			n3++
		case 1:
			n2++
		case 2:
			n1++
		}
	}

	for i := 0; i < 110; i++ {
		for j := 0; j < 110; j++ {
			for k := 0; k < 110; k++ {
				memo[i][j][k] = -1.0
			}
		}
	}
	memo[0][0][0] = 0.0
	out(rec(n3, n2, n1))
}
