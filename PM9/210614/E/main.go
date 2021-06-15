package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	p := make([][]int, N)
	pmax := make([]int, 5)
	for i := 0; i < N; i++ {
		p[i] = getInts(5)
		chmax(&pmax[0], p[i][0])
		chmax(&pmax[1], p[i][1])
		chmax(&pmax[2], p[i][2])
		chmax(&pmax[3], p[i][3])
		chmax(&pmax[4], p[i][4])
	}

	ans := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			maxP := make([]int, 5)
			for k := 0; k < 5; k++ {
				maxP[k] = max(p[i][k], p[j][k])
			}

			ans = max(ans, nmin(pmax[0], maxP[1], maxP[2], maxP[3], maxP[4]))
			ans = max(ans, nmin(pmax[1], maxP[0], maxP[2], maxP[3], maxP[4]))
			ans = max(ans, nmin(pmax[2], maxP[0], maxP[1], maxP[3], maxP[4]))
			ans = max(ans, nmin(pmax[3], maxP[0], maxP[1], maxP[2], maxP[4]))
			ans = max(ans, nmin(pmax[4], maxP[0], maxP[1], maxP[2], maxP[3]))
		}
	}
	out(ans)
}
