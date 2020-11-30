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

func move(x, y int, stat []byte, dir byte) (int, int, []byte) {
	flg := (abs(x) + abs(y)) % 2

	newStat := []byte{0, 0, 0}

	if flg == 0 { // △
		switch {
		case stat[0] == dir:
			x--
			newStat = []byte{stat[2], stat[0], stat[1]}
		case stat[1] == dir:
			x++
			newStat = []byte{stat[1], stat[2], stat[0]}
		case stat[2] == dir:
			y--
			newStat = []byte{stat[0], stat[1], stat[2]}
		}
	} else { // ▽
		switch {
		case stat[0] == dir:
			x--
			newStat = []byte{stat[2], stat[0], stat[1]}
		case stat[1] == dir:
			x++
			newStat = []byte{stat[1], stat[2], stat[0]}
		case stat[2] == dir:
			y++
			newStat = []byte{stat[0], stat[1], stat[2]}
		}
	}
	return x, y, newStat
}

type pos struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()
	x, y := 0, 0
	m := make(map[pos]bool)
	m[pos{x, y}] = true
	stat := []byte{'a', 'b', 'c'}
	for i := 0; i < len(s); i++ {
		x, y, stat = move(x, y, stat, s[i])
		// out(x, y, string(stat))
		m[pos{x, y}] = true
	}
	out(len(m))
}
