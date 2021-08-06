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
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = getInts(N - 1)
	}

	s := make([]pair, 0)
	used := make([]bool, N+1)
	for i := 1; i < N+1; i++ {
		x := a[i][0]
		if !used[i] && !used[x] && a[x][0] == i {
			used[i] = true
			used[x] = true
			s = append(s, pair{i, x})
		}
	}

	cnt := 0
	for {
		t := make([]pair, 0)
		for _, e := range s {
			i, j := e.x, e.y
			a[i] = a[i][1:]
			a[j] = a[j][1:]
		}
		used := make([]bool, N+1)
		for _, e := range s {
			i, j := e.x, e.y
			if len(a[i]) != 0 {
				x := a[i][0]
				if len(a[x]) != 0 && !used[i] && !used[x] && a[x][0] == i {
					used[i] = true
					used[x] = true
					t = append(t, pair{i, x})
				}
			}
			if len(a[j]) != 0 {
				x := a[j][0]
				if len(a[x]) != 0 && !used[j] && !used[x] && a[x][0] == j {
					used[j] = true
					used[x] = true
					t = append(t, pair{j, x})
				}
			}
		}
		cnt++
		s = t
		if len(t) == 0 {
			break
		}
	}

	ok := true
	for i := 1; i <= N; i++ {
		if len(a[i]) != 0 {
			ok = false
		}
	}

	if !ok {
		out(-1)
		return
	}
	out(cnt)
}
