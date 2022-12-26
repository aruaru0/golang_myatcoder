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

type pair struct {
	x, y int
}

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM : find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a / GCD(a, b) * b

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	e := make([]pair, K)
	for i := 0; i < K; i++ {
		x, y := getInt()-1, getInt()-1
		e[i] = pair{x, y}
	}
	c := make([]int, N)
	for i := 0; i < N; i++ {
		c[i] = i
	}
	for i := 0; i < K; i++ {
		x := e[i].x
		y := e[i].y
		c[x], c[y] = c[y], c[x]
	}
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}

	ans := 1
	for i := 0; i < N; i++ {
		cnt := 0
		p := i
		for {
			p = c[p]
			cnt++
			if p == i {
				break
			}
		}
		ans = LCM(ans, cnt)
	}
	out(ans)
}
