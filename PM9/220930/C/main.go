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

func match(sel_h, sel_w []int) bool {
	for h := 0; h < H2; h++ {
		for w := 0; w < W2; w++ {
			ah, aw := sel_h[h], sel_w[w]
			if a[ah][aw] != b[h][w] {
				return false
			}
		}
	}
	return true
}

var H1, H2 int
var W1, W2 int
var a [][]int
var b [][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H1, W1 = getI(), getI()
	a = make([][]int, H1)
	for i := 0; i < H1; i++ {
		a[i] = getInts(W1)
	}
	H2, W2 = getI(), getI()
	b = make([][]int, H2)
	for i := 0; i < H2; i++ {
		b[i] = getInts(W2)
	}

	BH := 1 << H1
	BW := 1 << W1
	for bh := 0; bh < BH; bh++ {
		sel_h := make([]int, 0)
		for i := 0; i < H1; i++ {
			if (bh>>i)%2 == 1 {
				sel_h = append(sel_h, i)
			}
		}
		if len(sel_h) != H2 {
			continue
		}
		for bw := 0; bw < BW; bw++ {
			sel_w := make([]int, 0)
			for i := 0; i < W1; i++ {
				if (bw>>i)%2 == 1 {
					sel_w = append(sel_w, i)
				}
			}
			if len(sel_w) != W2 {
				continue
			}
			if match(sel_h, sel_w) {
				out("Yes")
				return
			}
		}
	}
	out("No")
}
