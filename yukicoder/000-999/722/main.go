package main

import (
	"bufio"
	"fmt"
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
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	A, B := getS(), getS()

	flg := true
	s := 1
	if A[0] == '-' {
		s++
	}
	cnt := 0
	for i := s; i < len(A); i++ {
		if A[i] != '0' {
			flg = false
		}
		cnt++
	}
	if cnt < 2 {
		flg = false
	}
	s = 1
	if B[0] == '-' {
		s++
	}
	cnt = 0
	for i := s; i < len(B); i++ {
		if B[i] != '0' {
			flg = false
		}
		cnt++
	}
	if cnt < 2 {
		flg = false
	}
	a, _ := strconv.Atoi(A)
	b, _ := strconv.Atoi(B)

	if flg {
		out(a * b / 10)
		return
	}
	ans := a * b
	if ans > 99999999 || ans < -99999999 {
		out("E")
		return
	}
	out(ans)
}
