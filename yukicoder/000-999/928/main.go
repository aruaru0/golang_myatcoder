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
	P, Q, A := getI(), getI(), getI()

	l := 0
	r := int(1e9 + 1)
	for l+1 != r {
		m := (l + r) / 2
		x := ((100 + P) * m) / 100
		y := ((100+Q)*m)/100 + A
		if x >= y {
			r = m
		} else {
			l = m
		}
	}

	// 誤差のある小さい部分はカウント
	cnt := 0
	for i := 1; i < 1000; i++ {
		x := ((100 + P) * i) / 100
		y := ((100+Q)*i)/100 + A
		if x < y {
			cnt++
		}
	}

	// 探索周辺もカウントもチェック
	s := max(1000, l-100000)
	e := min(1e9, l+100000)
	cnt += s - 1000
	for i := s; i <= e; i++ {
		x := ((100 + P) * i) / 100
		y := ((100+Q)*i)/100 + A
		if x < y {
			cnt++
		} else {
			// out(i, x, y)
		}
	}

	out(cnt)
}
