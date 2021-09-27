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

type pos struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	a := make([][]int, 3)
	m := make(map[int]pos)
	for i := 0; i < 3; i++ {
		a[i] = getInts(3)
		for j := 0; j < 3; j++ {
			m[a[i][j]] = pos{i, j}
		}
	}

	N := getI()
	var t [3][3]bool
	for i := 0; i < N; i++ {
		b := getI()
		p, ok := m[b]
		if ok {
			t[p.x][p.y] = true
		}
	}

	ans := false
	for x := 0; x < 3; x++ {
		ok := true
		for y := 0; y < 3; y++ {
			ok = ok && t[x][y]
		}
		ans = ans || ok
	}
	for y := 0; y < 3; y++ {
		ok := true
		for x := 0; x < 3; x++ {
			ok = ok && t[x][y]
		}
		ans = ans || ok
	}
	if t[0][0] && t[1][1] && t[2][2] {
		ans = true
	}
	if t[2][0] && t[1][1] && t[0][2] {
		ans = true
	}

	if ans {
		out("Yes")
	} else {
		out("No")
	}
}
