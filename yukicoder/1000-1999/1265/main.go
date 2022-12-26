package main

import (
	"bufio"
	"fmt"
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
	idx0, idx1 int
	r          int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getI(), getI()
	}

	b := make([]pair, 0)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			dx := x[i] - x[j]
			dy := y[i] - y[j]
			b = append(b, pair{i, j, dx*dx + dy*dy})
		}
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i].r < b[j].r
	})

	m := make([]int, N)
	cnt := 0
	// out(b)
	for i := 0; i < len(b); i++ {
		if b[i].idx0 == 0 && m[b[i].idx1] == 0 {
			cnt++
			m[b[i].idx1] = 1
			// out(b[i])
		} else if m[b[i].idx0] == 0 && m[b[i].idx1] == 0 {
			m[b[i].idx0] = 1
			m[b[i].idx1] = 1
		}
	}
	out(cnt)
}
