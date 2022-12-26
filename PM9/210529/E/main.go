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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, N := getI(), getI(), getI()
	sr, sc := getI(), getI()
	s := getS()
	t := getS()

	l, r := sc, W-sc+1
	u, d := sr, H-sr+1

	tl, tr, tu, td := 0, 0, 0, 0
	al, ar, au, ad := 0, 0, 0, 0
	for i := 0; i < N; i++ {
		switch s[i] {
		case 'L':
			tl++
		case 'R':
			tr++
		case 'U':
			tu++
		case 'D':
			td++
		}
		// out("t", tl, tr, tu, td, "a", al, ar, au, ad)
		if tl-ar >= l || tr-al >= r || tu-ad >= u || td-au >= d {
			out("NO")
			return
		}

		switch t[i] {
		case 'L':
			if al-tr+1 < l {
				al++
			}
		case 'R':
			if ar-tl+1 < r {
				ar++
			}
		case 'U':
			if au-td+1 < u {
				au++
			}
		case 'D':
			if ad-tu+1 < d {
				ad++
			}
		}

	}

	out("YES")
}
