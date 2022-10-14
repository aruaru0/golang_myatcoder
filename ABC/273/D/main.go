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

const size = 210000

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, rs, cs := getI(), getI(), getI(), getI()
	h := make(map[int][]int)
	w := make(map[int][]int)
	N := getI()
	for i := 0; i < N; i++ {
		r, c := getI(), getI()
		h[r] = append(h[r], c)
		w[c] = append(w[c], r)
	}
	for i := range h {
		h[i] = append(h[i], 0, W+1)
		sort.Ints(h[i])
	}
	for i := range w {
		w[i] = append(w[i], 0, H+1)
		sort.Ints(w[i])
	}

	Q := getI()
	for q := 0; q < Q; q++ {
		d, l := getS(), getI()
		switch d[0] {
		case 'L':
			if len(h[rs]) == 0 {
				cs = max(1, cs-l)
			} else {
				pos := lowerBound(h[rs], cs)
				L := h[rs][pos-1]
				// out(d, l, ":", L, R, h[rs])
				if cs-l <= L {
					cs = L + 1
				} else {
					cs = cs - l
				}
			}
		case 'R':
			if len(h[rs]) == 0 {
				cs = min(W, cs+l)
			} else {
				pos := lowerBound(h[rs], cs)
				R := h[rs][pos]
				// out(d, l, ":", L, R, h[rs], pos)
				// out("new", cs, L, l, max(R-cs, l))
				if cs+l >= R {
					cs = R - 1
				} else {
					cs += l
				}
			}
		case 'U':
			if len(w[cs]) == 0 {
				rs = max(1, rs-l)
			} else {
				pos := lowerBound(w[cs], rs)
				L := w[cs][pos-1]
				// out(d, l, "rs", rs, ":", L, R, w[cs], pos)
				if rs-l <= L {
					rs = L + 1
				} else {
					rs = rs - l
				}
			}
		case 'D':
			if len(w[cs]) == 0 {
				rs = min(H, rs+l)
			} else {
				pos := lowerBound(w[cs], rs)
				R := w[cs][pos]
				// out(d, l, ":", L, R, w[cs])
				if rs+l >= R {
					rs = R - 1
				} else {
					rs += l
				}
			}
		}
		out(rs, cs)
	}
}
