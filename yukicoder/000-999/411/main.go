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

func f(a, b []int) bool {
	c := append(a, b...)
	cnt := 0
	for i := 1; i < len(c); i++ {
		if c[i] < c[i-1] {
			cnt++
		}
	}
	if cnt == 1 {
		return true
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()

	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i + 1
	}

	n := 1 << N
	cnt := 0
	for i := 0; i < n; i++ {
		a := make([]int, 0, N)
		b := make([]int, 0, N)
		for j := 0; j < N; j++ {
			if (i>>j)%2 == 0 {
				a = append(a, j+1)
			} else {
				b = append(b, j+1)
			}
		}
		if len(a) == 0 || len(b) == 0 {
			continue
		}
		if a[0] == K && f(a, b) == true {
			// out(a, b)
			cnt++
		}
	}
	out(cnt)
}
