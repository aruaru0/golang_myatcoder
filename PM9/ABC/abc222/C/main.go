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

type pair struct {
	win, num int
}

func check(x, y byte) (int, int) {
	if x == y {
		return 0, 0
	}
	if x == 'G' && y == 'C' {
		return 1, 0
	}
	if x == 'C' && y == 'P' {
		return 1, 0
	}
	if x == 'P' && y == 'G' {
		return 1, 0
	}
	return 0, 1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	N *= 2
	a := make([]string, N)
	for i := 0; i < N; i++ {
		a[i] = getS()
	}

	idx := make([]pair, N)
	for i := 0; i < N; i++ {
		idx[i] = pair{0, i}
	}

	for i := 0; i < M; i++ {

		for j := 0; j < N; j += 2 {
			p0, p1 := idx[j].num, idx[j+1].num
			x, y := check(a[p0][i], a[p1][i])
			idx[j].win += x
			idx[j+1].win += y
		}

		// sort
		sort.Slice(idx, func(i, j int) bool {
			if idx[i].win == idx[j].win {
				return idx[i].num < idx[j].num
			}
			return idx[i].win > idx[j].win
		})
		// for _, e := range idx {
		// 	out(e.num+1, e.win)
		// }
		// out("--------------")
	}

	for _, e := range idx {
		out(e.num + 1)
	}
}
