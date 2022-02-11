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

var di []int = []int{-1, 1, 0, 0}
var dj []int = []int{0, 0, -1, 1}

var n, k int
var ans int

func dfs(t [][]byte) {
	cnt := 0
	s := make([][]byte, n)
	// tをsにコピーしつつ、赤の個数(oの個数)を数える
	for i := 0; i < n; i++ {
		s[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			s[i][j] = t[i][j]
			if s[i][j] == 'o' {
				cnt++
			}
		}
	}
	// もし、赤の個数がkなら、ansをインクリメントし戻る（探索する必要はない）
	if cnt == k {
		ans++
		return
	}
	// そうでなければ、置ける場所を探索する
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] != '.' { // 置けない場所はスキップ
				continue
			}
			if cnt != 0 { // 既に赤がある場合、(i,j)の位置に赤を置いたとき、連続になるか調べる
				ok := false
				for v := 0; v < 4; v++ {
					ni := i + di[v]
					nj := j + dj[v]
					if ni < 0 || nj < 0 || ni >= n || nj >= n {
						continue
					}
					if s[ni][nj] == 'o' {
						ok = true
					}
				}
				// 連続にならないなら、ここには置けない
				if !ok {
					continue
				}
			}
			// 赤に塗って探索
			s[i][j] = 'o'
			dfs(s)
			// 白でない色にして探索（例えば黒）※探索済みとマークする
			s[i][j] = '#'
			dfs(s)
			return
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, k = getI(), getI()
	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		s[i] = []byte(getS())
	}
	dfs(s)
	out(ans)
}
