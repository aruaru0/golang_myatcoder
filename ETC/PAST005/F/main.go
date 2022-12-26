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
	N, M := getI(), getI()
	a := make([]int, M)
	b := make([]int, M)
	c := make([]int, M)
	for i := 0; i < M; i++ {
		a[i], b[i], c[i] = getI()-1, getI()-1, getI()-1
	}
	n := 1 << N
	tot := 0
	for i := 0; i < n; i++ {
		x := make([]int, N)
		for j := 0; j < N; j++ {
			if (i>>j)&1 == 1 {
				x[j] = 1
			}
		}
		ng := false
		ok := 0
		tt := make([]bool, N)
		for j := 0; j < M; j++ {
			cnt := x[a[j]] + x[b[j]] + x[c[j]]
			if cnt == 3 {
				ng = true
				break
			}
			if cnt == 2 {
				target := 0
				if x[a[j]] == 0 {
					target = a[j]
				}
				if x[b[j]] == 0 {
					target = b[j]
				}
				if x[c[j]] == 0 {
					target = c[j]
				}
				if tt[target] == false {
					ok++
					tt[target] = true
				}
			}
		}
		if ng == false {
			// out(x, ok)
			tot = max(tot, ok)
		}
	}
	out(tot)
}
