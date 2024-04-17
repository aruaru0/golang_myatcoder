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

const (
	X             = 1e6 + 1
	MIN_LETTER    = 'a'
	DELTA_LETTERS = 'A' - 'a'
)

var (
	grid      [3][3]int
	positions [3][3]int
)

func f(x int) int {
	if x%2 != 0 {
		return 1
	}
	return 2
}

func checkDir(x, y, dx, dy int) bool {
	// 各方向に3マス揃っているか計算
	now := positions[x][y]
	if now == 0 {
		return false
	}
	for i := 0; i < 3; i++ {
		if now != positions[x+i*dx][y+i*dy] {
			return false
		}
	}
	return true
}

func check() bool {
	if checkDir(0, 0, 0, 1) || checkDir(0, 0, 1, 0) || checkDir(0, 0, 1, 1) {
		return true
	}
	if checkDir(1, 0, 0, 1) || checkDir(2, 0, 0, 1) {
		return true
	}
	if checkDir(0, 1, 1, 0) || checkDir(0, 2, 1, 0) {
		return true
	}
	if checkDir(2, 0, -1, 1) {
		return true
	}
	return false
}

func solve(x int) bool {
	if x == 10 { // ９マス全て埋まっていれば、点数計算
		takScore := 0
		aokScore := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if positions[i][j] == 1 {
					takScore += grid[i][j]
				} else {
					aokScore += grid[i][j]
				}
			}
		}
		if aokScore > takScore {
			return true
		}
		return false
	}
	// 打って縦横斜めのどこかに揃うか計算
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if positions[i][j] == 0 {
				positions[i][j] = f(x)
				if check() { // 揃えば勝ち
					positions[i][j] = 0
					return true
				}
				positions[i][j] = 0
			}
		}
	}
	// 全てのパターンについて検査
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if positions[i][j] == 0 {
				positions[i][j] = f(x)
				if !solve(x + 1) { // 結果として負けていなければ先行の価値
					positions[i][j] = 0
					return true
				}
				positions[i][j] = 0
			}
		}
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[i][j] = getI()
		}
	}
	if solve(1) {
		out("Takahashi")
	} else {
		out("Aoki")
	}
}
