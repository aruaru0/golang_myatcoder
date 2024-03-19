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

var N, H, W int
var a, b []int

func disp(p [][]bool) {
	out("----------")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if p[i][j] == true {
				fmt.Fprint(wr, "#")
			} else {
				fmt.Fprint(wr, ".")
			}
		}
		out()
	}
}

func makeP() [][]bool {
	ret := make([][]bool, H)
	for i := 0; i < H; i++ {
		ret[i] = make([]bool, W)
	}
	return ret
}

func copyP(p [][]bool) [][]bool {
	ret := make([][]bool, H)
	for i := 0; i < H; i++ {
		ret[i] = make([]bool, W)
		for j := 0; j < W; j++ {
			ret[i][j] = p[i][j]
		}
	}
	return ret
}

func rec(p [][]bool, s []int) bool {
	sx, sy := -1, -1
	ok := true
	// 空白があるかどうかチェック
LOOP:
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if p[i][j] == false {
				// 埋まっていない場所があれば、そこを返す
				sx, sy = j, i
				ok = false
				break LOOP
			}
		}
	}
	// もし全部埋まっていたらOK
	if ok {
		return true
	}

	// もし、埋まっていないのにパーツがなければfalse
	if len(s) == 0 {
		return false
	}

	n := s[0]
	h, w := a[n], b[n]

	for l := 0; l < 2; l++ {

		ok := true
		// パーツが置けるかどうかチェック
	LOOP2:
		for i := sy; i < sy+h; i++ {
			for j := sx; j < sx+w; j++ {
				// はみ出したらOUT
				if i >= H || j >= W {
					ok = false
					break LOOP2
				}
				// すでにおいていたらOUT
				if p[i][j] {
					ok = false
					break LOOP2
				}
			}
		}

		if ok {
			q := copyP(p)
			for i := sy; i < sy+h; i++ {
				for j := sx; j < sx+w; j++ {
					q[i][j] = true
				}
			}
			if rec(q, s[1:]) {
				return true
			}
		}

		h, w = w, h
	}

	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, H, W = getI(), getI(), getI()
	a = make([]int, N)
	b = make([]int, N)
	for i := 0; i < N; i++ {
		a[i], b[i] = getI(), getI()
	}

	s := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = i
	}

	for {
		p := makeP()
		if rec(p, s) {
			out("Yes")
			return
		}
		if NextPermutation(sort.IntSlice(s)) == false {
			break
		}
	}

	out("No")
}
