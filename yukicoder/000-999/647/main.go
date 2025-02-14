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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getI(), getI()
	}
	M := getI()
	x := make([]int, M)
	y := make([]int, M)
	buy := make([]int, M)
	for i := 0; i < M; i++ {
		x[i], y[i] = getI(), getI()
	}

	for i := 0; i < N; i++ {
		cost := a[i]
		taste := b[i]
		for j := 0; j < M; j++ {
			if x[j] <= cost && y[j] >= taste {
				buy[j]++
			}
		}
	}

	mx := 0
	pos := make([]int, 0)
	for i := 0; i < M; i++ {
		if mx < buy[i] {
			pos = make([]int, 0, M)
			mx = buy[i]
			pos = append(pos, i)
		} else if mx == buy[i] {
			pos = append(pos, i)
		}
	}
	if mx == 0 {
		out(0)
		return
	}
	for _, e := range pos {
		out(e + 1)
	}
}
