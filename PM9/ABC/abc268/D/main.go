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
	N, M := getI(), getI()
	ss := make([]string, N)
	l := 0
	for i := 0; i < N; i++ {
		ss[i] = getS()
		l += len(ss[i])
	}
	mm := make(map[string]bool)
	for i := 0; i < M; i++ {
		mm[getS()] = true
	}

	// もし１文字列の場合、条件を満たせばOK、でなければNG
	if N == 1 {
		s := ss[0]
		if mm[s] || len(s) < 3 || len(s) > 16 {
			out(-1)
		} else {
			out(s)
		}
		return
	}

	// S1..S8で１６文字を超えたらNG
	if l > 16 {
		out(-1)
		return
	}

	// 順列を作るテーブル
	sel := make([]int, N)
	for i := 0; i < N; i++ {
		sel[i] = i
	}

	for {
		// ありうるパターンを列挙
		st := []string{""}
		g := l
		for i, v := range sel {
			w := ss[v]
			g -= len(w)
			nst := []string{}
			for _, pre := range st {
				s := pre + w
				if i+1 == N {
					if len(s) > 3 && len(s) <= 16 && !mm[s] {
						out(s)
						return
					}
				} else {
					// 文字列のパターンをすべて作成する
					for len(s)+g < 16 {
						s += "_"
						nst = append(nst, s)
					}
				}
			}
			st = nst
		}

		// 順列で組み合わせを列挙する
		if NextPermutation(sort.IntSlice(sel)) == false {
			break
		}
	}

	out(-1)
}
