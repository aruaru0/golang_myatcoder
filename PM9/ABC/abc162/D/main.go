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

func rgb2idx(c byte) int {
	if c == 'R' {
		return 0
	}
	if c == 'G' {
		return 1
	}
	return 2
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	s := getS()
	rgb := make([][3]int, N)
	for i := 0; i < N; i++ {
		idx := rgb2idx(s[i])
		rgb[i][idx]++
	}

	for i := N - 1; i > 0; i-- {
		for j := 0; j < 3; j++ {
			rgb[i-1][j] += rgb[i][j]
		}
	}

	tot := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if s[i] == s[j] {
				continue
			}
			flg := make([]bool, 3)
			flg[rgb2idx(s[i])] = true
			flg[rgb2idx(s[j])] = true
			pos := 0
			for n := 0; n < 3; n++ {
				if flg[n] == false {
					tot += rgb[j][n]
					pos = n
				}
			}
			k := j + j - i
			if k < N && rgb2idx(s[k]) == pos {
				tot--
			}
		}
	}
	out(tot)
}
