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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type B struct {
	l, x int
	s    byte
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	w := getInts(n)

	bs := make([]B, 0)
	bs = append(bs, B{n, 0, 'L'})
	bs = append(bs, B{0, 0, 'L'})

	// 区間を返す
	get := func(b B, i int) (int, int) {
		if b.s == 'L' {
			return b.x, b.x + w[i]
		}
		return b.x - w[i], b.x
	}

	q := getI()
	for qi := 0; qi < q; qi++ {
		t, x := getI(), getI()
		if t == 3 {
			ac, wa := 0, len(bs)
			for ac+1 < wa {
				wj := (ac + wa) / 2
				l, r := get(bs[wj], bs[wj].l)
				if l <= x && x < r {
					ac = wj
				} else {
					wa = wj
				}
			}
			ans := n
			if wa < len(bs) {
				bi := wa
				ac = bs[bi].l
				wa = bs[bi-1].l
				for ac+1 < wa {
					wj := (ac + wa) / 2
					l, r := get(bs[bi], wj)
					if l <= x && x < r {
						wa = wj
					} else {
						ac = wj
					}
				}
				ans = n - wa
			}
			out(ans)
		} else {
			x--
			if x == 0 {
				continue
			}

			for {
				b := bs[len(bs)-1]
				bs = bs[:len(bs)-1]
				if x < bs[len(bs)-1].l {
					bs = append(bs, B{x, b.x, b.s})
					break
				}
			}
			{
				b := bs[len(bs)-1]
				l, r := get(b, b.l)
				if t == 1 {
					bs = append(bs, B{0, l, 'L'})
				} else {
					bs = append(bs, B{0, r, 'R'})
				}
			}
		}
	}

}
