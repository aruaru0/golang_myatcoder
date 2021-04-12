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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := make([]string, 3)
	s[0] = getS()
	s[1] = getS()
	s[2] = getS()

	m := make(map[byte]int)
	for i := 0; i < 3; i++ {
		for _, e := range s[i] {
			m[byte(e)]++
		}
	}
	if len(m) > 10 {
		out("UNSOLVABLE")
		return
	}

	t := make([]byte, 0)
	for e := range m {
		t = append(t, e)
	}
	for i, e := range t {
		m[e] = i
	}
	st := make([][]int, 3)
	for i := 0; i < 3; i++ {
		for _, e := range s[i] {
			st[i] = append(st[i], m[byte(e)])
		}
	}

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for {
		if a[st[0][0]] == 0 || a[st[1][0]] == 0 || a[st[2][0]] == 0 {
		} else {
			x := make([]int, 3)
			for i := 0; i < 3; i++ {
				tot := 0
				for _, e := range st[i] {
					tot = tot*10 + a[e]
				}
				x[i] = tot
			}
			if x[0]+x[1] == x[2] {
				out(x[0])
				out(x[1])
				out(x[2])
				return
			}

		}

		if !NextPermutation(sort.IntSlice(a)) {
			break
		}
	}
	out("UNSOLVABLE")
}
