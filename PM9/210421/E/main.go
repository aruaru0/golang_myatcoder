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

//
// バグ修正（copyを忘れていた)
// 問題なければこちらの組み合わせ列挙を使う
//
func generateComb(index []int, s, r int, ch chan []int) {
	if r != 0 {
		if s < 0 {
			return
		}
		generateComb(index, s-1, r, ch)
		index[r-1] = s
		generateComb(index, s-1, r-1, ch)
	} else {
		out := make([]int, len(index))
		copy(out, index)
		ch <- out
	}
	return
}

func foreachComb(n, k int, ch chan []int) {
	index := make([]int, k)
	generateComb(index, n-1, k, ch)
	close(ch)
}

func f(x, y, z int) int {
	tot := 0
	for i := 0; i < C; i++ {
		tot += p[0][i] * d[i][x]
	}
	for i := 0; i < C; i++ {
		tot += p[1][i] * d[i][y]
	}
	for i := 0; i < C; i++ {
		tot += p[2][i] * d[i][z]
	}
	return tot
}

var N, C int
var p [][]int
var d [][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, C = getI(), getI()
	d = make([][]int, C)
	for i := 0; i < C; i++ {
		d[i] = getInts(C)
	}
	c := make([][]int, N)
	for i := 0; i < N; i++ {
		c[i] = getInts(N)
	}

	p = make([][]int, 3)
	for i := 0; i < 3; i++ {
		p[i] = make([]int, C)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			pos := (i + j) % 3
			// out(pos, c[i][j])
			p[pos][c[i][j]-1]++
		}
	}

	// for i := 0; i < 3; i++ {
	// 	out(p[i])
	// }

	ans := int(1e18)
	for i := 0; i < C; i++ {
		for j := 0; j < C; j++ {
			if i == j {
				continue
			}
			for k := 0; k < C; k++ {
				if i == k || j == k {
					continue
				}

				ret := f(i, j, k)
				ans = min(ans, ret)
				// out(i+1, j+1, k+1, ret)
			}
		}
	}
	out(ans)
}
