package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	x := getInts(N)
	a := getInts(N)

	xn := make([][]int, 100)
	for i := 0; i < 100; i++ {
		xn[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		xn[0][i] = i
	}
	for i := 0; i < N; i++ {
		xn[1][i] = xn[0][x[i]-1]
	}

	for i := 2; i < 100; i++ {
		for j := 0; j < N; j++ {
			xn[i][j] = xn[i-1][xn[i-1][j]]
		}
	}

	pos := 1
	for K > 0 {
		if K%2 == 1 {
			tmp := make([]int, N)
			// out("use", pos, xn[pos])
			for i := 0; i < N; i++ {
				tmp[i] = a[xn[pos][i]]
			}
			a = tmp
		}
		pos++
		K /= 2
	}

	for _, e := range a {
		fmt.Fprint(wr, e, " ")
	}
	out()

	// for i := 0; i < 6; i++ {
	// 	out(xn[i])
	// }

	// out("----")
	// for i := 0; i < N; i++ {
	// 	a[i] = i
	// }
	// out(a)
	// for i := 0; i < 8; i++ {
	// 	tmp := make([]int, N)
	// 	for j := 0; j < N; j++ {
	// 		tmp[j] = a[x[j]-1]
	// 	}
	// 	a = tmp
	// 	out(i+1, a, x)
	// }

	// out("----")
	// out(a, K)
}
