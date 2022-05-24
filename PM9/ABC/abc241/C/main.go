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

func calc(dx, dy []int, N int, s []string) bool {
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			cnt := 0
			for i := 0; i < 6; i++ {
				px := x + dx[i]
				py := y + dy[i]
				if px < 0 || px >= N || py < 0 || py >= N {
					cnt = 0
					break
				}
				if s[py][px] == '#' {
					cnt++
				}
			}
			if cnt >= 4 {
				return true
			}
		}
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
	}

	dx := []int{0, 1, 2, 3, 4, 5}
	dy := []int{0, 0, 0, 0, 0, 0}
	if calc(dx, dy, N, s) {
		out("Yes")
		return
	}
	dx = []int{0, 0, 0, 0, 0, 0}
	dy = []int{0, 1, 2, 3, 4, 5}
	if calc(dx, dy, N, s) {
		out("Yes")
		return
	}

	dx = []int{0, 1, 2, 3, 4, 5}
	dy = []int{0, 1, 2, 3, 4, 5}
	if calc(dx, dy, N, s) {
		out("Yes")
		return
	}

	dx = []int{5, 4, 3, 2, 1, 0}
	dy = []int{0, 1, 2, 3, 4, 5}
	if calc(dx, dy, N, s) {
		out("Yes")
		return
	}

	out("No")
}
