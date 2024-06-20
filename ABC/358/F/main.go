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
	n, m, k := getI(), getI(), getI()

	if k < n || (k-n)%2 == 1 {
		out("No")
		return
	}

	s := make([][]byte, n*2+1)
	for i := 0; i < n*2+1; i++ {
		s[i] = make([]byte, m*2+1)
		for j := 0; j <= m*2; j++ {
			s[i][j] = '+'
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s[i*2+1][j*2+1] = 'o'
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			s[i*2+1][j*2+2] = '|'
		}
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			s[i*2+2][j*2+1] = '-'
		}
	}
	s[0][m*2-1] = 'S'
	s[n*2][m*2-1] = 'G'

	for i := 0; i < n-1; i++ {
		s[i*2+2][m*2-1] = '.'
	}

	k -= n
	for i := 0; i < n-1; i += 2 {
		for j := m - 1; j >= 1; j-- {
			if k > 0 {
				s[i*2+2][j*2+1] = '-'
				s[i*2+2][j*2-1] = '.'
				s[i*2+1][j*2] = '.'
				s[i*2+3][j*2] = '.'
				k -= 2
			}
		}
	}
	if n%2 == 1 {
		for j := 0; j < m-2; j += 2 {
			if k > 0 {
				s[n*2-3][j*2+2] = '|'
				s[n*2-1][j*2+2] = '.'
				s[n*2-2][j*2+1] = '.'
				s[n*2-2][j*2+3] = '.'
				k -= 2
			}
		}
	}

	out("Yes")
	for i := 0; i < n*2+1; i++ {
		out(string(s[i]))
	}
}
