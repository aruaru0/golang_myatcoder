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
	N := getI()
	C := make([][]int, N)
	for i := 0; i < N; i++ {
		C[i] = getInts(N)
	}

	if N == 1 {
		out("Yes")
		out(C[0][0], 0)
		return
	}

	ok := true
L0:
	for r := 1; r < N; r++ {
		v := C[r][0] - C[r-1][0]
		for c := 0; c < N; c++ {
			if C[r][c]-C[r-1][c] != v {
				ok = false
				break L0
			}
		}
	}
L1:
	for c := 1; c < N; c++ {
		v := C[0][c] - C[0][c-1]
		for r := 0; r < N; r++ {
			if C[r][c]-C[r][c-1] != v {
				ok = false
				break L1
			}
		}
	}
	if !ok {
		out("No")
		return
	}

	out("Yes")
	a := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		a[i] = C[i][0] - C[i+1][0]
	}
	b := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		b[i] = C[0][i] - C[0][i+1]
	}
	A := a[0]
	// B := b[0]
	for i := 1; i < len(a); i++ {
		a[i] += a[i-1]
		b[i] += b[i-1]
		A = max(A, a[i])
		//B = max(B, b[i])
	}
	A = max(A, 0)
	fmt.Fprint(wr, A, " ")
	for i := 0; i < N-1; i++ {
		fmt.Fprint(wr, A-a[i], " ")
	}
	out()
	B := C[0][0] - A
	fmt.Fprint(wr, B, " ")
	for i := 0; i < N-1; i++ {
		fmt.Fprint(wr, B-b[i], " ")
	}
	out()
}
