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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	a := getInts(n)
	b := getInts(n)

	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
		for j := 0; j < n; j++ {
			c[i][j] = a[i] * b[j] % m
		}
	}

	w := n * 3
	d := make([][]int, w)
	for i := 0; i < w; i++ {
		d[i] = make([]int, w)
	}
	{
		e := make([][]int, w)
		for i := 0; i < w; i++ {
			e[i] = make([]int, w)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				e[i+2][j+2] -= c[i][j]
				e[i+n*2][j+n*2] += c[i][j]
			}
		}
		for i := 0; i < w-1; i++ {
			for j := 0; j < w-1; j++ {
				e[i+1][j+1] += e[i][j]
			}
		}
		for i := 0; i < w; i++ {
			for j := 0; j < w; j++ {
				d[i][j] += e[i][j]
			}
		}
	}

	{
		e := make([][]int, w)
		for i := 0; i < w; i++ {
			e[i] = make([]int, w)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				e[i+2][j+n*2-1] += c[i][j]
				e[i+n*2][j+1] -= c[i][j]
			}
		}
		for i := 0; i < w-1; i++ {
			for j := 0; j < w-1; j++ {
				e[i+1][j] += e[i][j+1]
			}
		}
		for i := 0; i < w; i++ {
			for j := 0; j < w; j++ {
				d[i][j] += e[i][j]
			}
		}
	}

	{
		e := make([][]int, w)
		for i := 0; i < w; i++ {
			e[i] = make([]int, w)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				{
					d[i+1][j+1] += c[i][j] * (n - 1)
					d[i+1][j+n*2] -= c[i][j] * (n - 1)
					d[i+n*2][j+1] -= c[i][j] * (n - 1)
					d[i+n*2][j+n*2] += c[i][j] * (n - 1)
				}
			}
		}
	}

	for i := 0; i < w-1; i++ {
		for j := 0; j < w-1; j++ {
			d[i+1][j] += d[i][j]
		}
	}
	for i := 0; i < w-1; i++ {
		for j := 0; j < w-1; j++ {
			d[i][j+1] += d[i][j]
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans ^= d[n+i][n+j] + n*i + j
		}
	}
	out(ans)
}
