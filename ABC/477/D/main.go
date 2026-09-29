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

type interval struct {
	l, r int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, 1024*1024*64)

	N, Q := getI(), getI()

	// 各マスのタイプ 1 クエリ時刻を記録
	pos := make([][]int, N)
	// タイプ 2 クエリの時刻と色を記録
	type2idx := make([]int, 0, Q)
	type2col := make([]byte, Q)

	for qi := 0; qi < Q; qi++ {
		t := getI()
		if t == 1 {
			x := getI() - 1
			pos[x] = append(pos[x], qi)
		} else {
			c := getS()[0]
			type2idx = append(type2idx, qi)
			type2col[qi] = c
		}
	}

	ans := make([]byte, N)

	for i := 0; i < N; i++ {
		// マス i にタイルが存在しなかった区間 [L, R) を構築
		intervals := make([]interval, 0, len(pos[i])/2+1)
		m := len(pos[i])
		if m == 0 {
			intervals = append(intervals, interval{0, Q})
		} else {
			intervals = append(intervals, interval{0, pos[i][0]})
			for j := 1; j+1 < m; j += 2 {
				intervals = append(intervals, interval{pos[i][j], pos[i][j+1]})
			}
			if m%2 == 0 {
				intervals = append(intervals, interval{pos[i][m-1], Q})
			}
		}

		// 時系列の後ろから区間を走査し、最新のタイプ 2 クエリを二分探索
		col := byte('a')
		for k := len(intervals) - 1; k >= 0; k-- {
			L, R := intervals[k].l, intervals[k].r
			if L >= R {
				continue
			}

			// type2Indices の中で R 未満の最大要素のインデックスを探す
			idx := sort.Search(len(type2idx), func(x int) bool {
				return type2idx[x] >= R
			})

			// idx - 1 が R 未満で最大の要素を指す
			if idx > 0 && type2idx[idx-1] >= L {
				col = type2col[type2idx[idx-1]]
				break
			}
		}
		ans[i] = col
	}

	out(string(ans))
}
