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
	a := make([][]int, N)
	col := make([]int, 0)
	row := make([]int, 0)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
		for j := 0; j < N; j++ {
			a[i][j] = -1
		}
	}
	for i := 0; i < N; i++ {
		s := getI()
		if s == 0 {
			for j := 0; j < N; j++ {
				a[i][j] = 0
			}
		} else if s == 2 {
			for j := 0; j < N; j++ {
				a[i][j] = 1
			}
		} else {
			col = append(col, i)
		}
	}

	for i := 0; i < N; i++ {
		t := getI()
		if t == 0 {
			for j := 0; j < N; j++ {
				a[j][i] = 0
			}
		} else if t == 2 {
			for j := 0; j < N; j++ {
				a[j][i] = 1
			}
		} else {
			row = append(row, i)
		}
	}

	cnt := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if a[i][j] == 1 {
				cnt++
			}
		}
	}

	// next
	c := 0
	for _, e := range col {
		ok := false
		for i := 0; i < N; i++ {
			if a[e][i] == 1 {
				ok = true
			}
		}
		if !ok {
			c++
		}
	}
	r := 0
	for _, e := range row {
		ok := false
		for i := 0; i < N; i++ {
			if a[i][e] == 1 {
				ok = true
			}
		}
		if !ok {
			r++
		}
	}

	// out(a)
	// out(col, row)
	// out(c, r)
	out(cnt + max(c, r))
}
