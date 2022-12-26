package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func getf() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

var N int
var Pa, Pb int
var a, b []int

func f() int {
	x := make([]int, N)
	y := make([]int, N)
	copy(x, a)
	copy(y, b)

	sumA := 0
	sumB := 0

	for i := 0; i < N; i++ {
		var A, B int
		rx := rand.Intn(1000)
		if i == N-1 || rx < Pa {
			A = x[0]
			x = x[1:]
		} else {
			n := rand.Intn(len(x)-1) + 1
			A = x[n]
			x = append(x[0:n], x[n+1:]...)
		}
		ry := rand.Intn(1000)
		if i == N-1 || ry < Pb {
			B = y[0]
			y = y[1:]
		} else {
			n := rand.Intn(len(y)-1) + 1
			B = y[n]
			y = append(y[0:n], y[n+1:]...)
		}
		// out(A, B)
		if A > B {
			sumA += A + B
		} else if A < B {
			sumB += A + B
		}
	}
	return sumA - sumB
}

func main() {
	sc.Split(bufio.ScanWords)
	N = getInt()
	x, y := getf(), getf()
	a = getInts(N)
	b = getInts(N)

	Pa = int(x * 1000)
	Pb = int(y * 1000)

	sort.Ints(a)
	sort.Ints(b)

	win := 0.0
	tot := 0.0
	for i := 0; i < 200000; i++ {
		res := f()
		if res > 0 {
			win++
		}
		tot++
	}
	out(win / tot)
}
