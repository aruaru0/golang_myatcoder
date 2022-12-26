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
	H, W := getI(), getI()
	s := make([]string, H)
	x := make([][]int, H)
	y := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
		x[i] = make([]int, W)
		y[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		cnt := 1
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				cnt = 1
				continue
			}
			x[i][j] += cnt
			cnt++
		}
		for j := W - 2; j >= 0; j-- {
			if s[i][j] == '#' {
				continue
			}
			if s[i][j+1] != '#' {
				x[i][j] = x[i][j+1]
			}
		}
	}

	for j := 0; j < W; j++ {
		cnt := 1
		for i := 0; i < H; i++ {
			if s[i][j] == '#' {
				cnt = 1
				continue
			}
			y[i][j] += cnt
			cnt++
		}
		for i := H - 2; i >= 0; i-- {
			if s[i][j] == '#' {
				continue
			}
			if s[i+1][j] != '#' {
				y[i][j] = y[i+1][j]
			}
		}
	}

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				continue
			}
			ans = max(ans, x[i][j]+y[i][j]-1)
		}
	}
	out(ans)
}
