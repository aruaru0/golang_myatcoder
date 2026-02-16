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

type pair struct {
	x, y, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, N := getI(), getI(), getI()

	h := make([]pair, N)
	w := make([]pair, N)
	for i := 0; i < N; i++ {
		x, y := getI(), getI()
		h[i] = pair{x, y, i}
		w[i] = pair{x, y, i}
	}

	sort.Slice(h, func(i, j int) bool {
		return h[i].x > h[j].x
	})
	sort.Slice(w, func(i, j int) bool {
		return w[i].y > w[j].y
	})

	used := make(map[int]bool)

	posH := make([]int, N)
	posW := make([]int, N)

	ph, pw := 0, 0
	for i := 0; i < N; i++ {
		// すでに使ったやつを削除
		for len(h) > 0 && used[h[0].idx] == true {
			h = h[1:]
		}
		for len(w) > 0 && used[w[0].idx] == true {
			w = w[1:]
		}
		// out(H, W)
		if H == h[0].x { // 幅が現在のサイズと等しい時
			// out("H", h[0], ph, pw)
			posH[h[0].idx] = ph
			posW[h[0].idx] = pw
			pw += h[0].y
			used[h[0].idx] = true
			W -= h[0].y

		} else if W == w[0].y { //
			// out("W", w[0], ph, pw)
			posH[w[0].idx] = ph
			posW[w[0].idx] = pw
			ph += w[0].x
			used[w[0].idx] = true
			H -= w[0].x
		}
	}

	for i := 0; i < N; i++ {
		out(posH[i]+1, posW[i]+1)
	}

}
