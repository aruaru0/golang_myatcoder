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
	y, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, W := getI(), getI()

	X := make([]int, N)
	Y := make([]int, N)
	blocks := make([][]struct{ y, index int }, W+1)

	for i := 0; i < N; i++ {
		X[i], Y[i] = getI(), getI()
		blocks[X[i]] = append(blocks[X[i]], struct{ y, index int }{Y[i], i})
	}

	cnt := make([]int, N)
	disappear := make([]int, N+1)
	for i := range disappear {
		disappear[i] = -1
	}

	for x := 1; x <= W; x++ {
		sort.Slice(blocks[x], func(i, j int) bool {
			return blocks[x][i].y < blocks[x][j].y
		})
		for j, block := range blocks[x] {
			cnt[block.index] = j
			disappear[j] = max(disappear[j], block.y)
		}
		disappear[len(blocks[x])] = 1e10
	}

	for i := 0; i < N; i++ {
		disappear[i+1] = max(disappear[i+1], disappear[i]+1)
	}

	Q := getI()
	for i := 0; i < Q; i++ {
		T, A := getI(), getI()
		A--
		if T < disappear[cnt[A]] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
