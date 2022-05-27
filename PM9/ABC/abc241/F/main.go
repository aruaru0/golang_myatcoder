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
	_, _, n := getI(), getI(), getI()
	sx, sy, gx, gy := getI(), getI(), getI(), getI()

	rows := make(map[int][]int)
	cols := make(map[int][]int)
	blocks := make(map[[2]int]bool)

	for i := 0; i < n; i++ {
		x, y := getI(), getI()
		cols[x] = append(cols[x], y)
		rows[y] = append(rows[y], x)
		blocks[[2]int{x, y}] = true
	}

	// BSF
	que := make([][2]int, 0)
	que = append(que, [2]int{sx, sy})

	dist := make(map[[2]int]int)
	// pos :=
	dist[[2]int{sx, sy}] = 1
	const inf = int(1e18)

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]

		row, col := cur[0], cur[1]
		// stop if goal
		if row == gx && col == gy {
			break
		}

		// search next block colL, colR
		nps := make([][2]int, 0)
		colL := -inf
		colR := inf
		for _, c := range cols[row] {
			if c < col {
				colL = max(colL, c+1)
			} else {
				colR = min(colR, c-1)
			}
		}
		if colL != -inf {
			nps = append(nps, [2]int{row, colL})
		}
		if colR != inf {
			nps = append(nps, [2]int{row, colR})
		}
		// search next block rowU, rowD
		rowT := -inf
		rowB := inf
		for _, r := range rows[col] {
			if r < row {
				rowT = max(rowT, r+1)
			} else {
				rowB = min(rowB, r-1)
			}
		}
		if rowT != -inf {
			nps = append(nps, [2]int{rowT, col})
		}
		if rowB != inf {
			nps = append(nps, [2]int{rowB, col})
		}

		// seach next node
		for _, e := range nps {
			_, ok := dist[e]
			// if already visited or '#' block -> skip
			if ok || blocks[e] {
				continue
			}
			// set dist and push to queue
			dist[e] = dist[cur] + 1
			que = append(que, e)
		}
	}
	ans := dist[[2]int{gx, gy}] - 1
	out(ans)
}
