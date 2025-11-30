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

type pos struct {
	i, j int
}

func solve() {
	h, w := getI(), getI()
	x := getInts(h)
	y := getInts(w)

	rid := make(map[int]int)
	cid := make(map[int]int)

	for i := 0; i < h; i++ {
		rid[x[i]] = i
	}
	for i := 0; i < w; i++ {
		cid[y[i]] = i
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})

	sort.Slice(y, func(i, j int) bool {
		return y[i] > y[j]
	})

	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}

	{
		xi, yi := 0, 0
		cand := make([]pos, 0)
		for v := h * w; v >= 1; v-- {
			if xi < h && x[xi] == v {
				for j := 0; j < yi; j++ {
					cand = append(cand, pos{xi, j})
				}
				xi++
			}
			if yi < w && y[yi] == v {
				for i := 0; i < xi; i++ {
					cand = append(cand, pos{i, yi})
				}
				yi++
			}
			if len(cand) == 0 {
				out("No")
				return
			}
			cur := cand[len(cand)-1]
			cand = cand[:len(cand)-1]
			a[cur.i][cur.j] = v
		}
	}

	ans := make([][]int, h)
	for i := 0; i < h; i++ {
		ans[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans[rid[x[i]]][cid[y[j]]] = a[i][j]
		}
	}
	out("Yes")
	for i := 0; i < h; i++ {
		outSlice(ans[i])
	}

}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	T := getI()
	for i := 0; i < T; i++ {
		solve()
	}
}
