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

const INF = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	h, w := getI(), getI()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = getInts(w)
	}
	var rows [][2]int

	//行に関する条件の評価
	for i := 0; i < h; i++ {
		Min, Max := INF, -INF
		for j := 0; j < w; j++ {
			if a[i][j] == 0 {
				continue
			}
			Min = min(Min, a[i][j])
			Max = max(Max, a[i][j])
		}
		if Min <= Max {
			rows = append(rows, [2]int{Min, Max})
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i][0] == rows[j][0] {
			return rows[i][1] < rows[j][1]
		}
		return rows[i][0] < rows[j][0]
	})
	for i := range rows {
		if i == len(rows)-1 {
			break
		}
		if rows[i][1] > rows[i+1][0] {
			out("No")
			return
		}
	}

	g := make([][]int, w+h*w+1)
	//fmt.Println(len(g))
	for r := 0; r < h; r++ {
		var cols [][2]int
		for c := 0; c < w; c++ {
			if a[r][c] == 0 {
				continue
			}
			cols = append(cols, [2]int{a[r][c], c})
		}
		sort.Slice(cols, func(i, j int) bool {
			if cols[i][0] == cols[j][0] {
				return cols[i][1] < cols[j][1]
			}
			return cols[i][0] < cols[j][0]
		})
		for i := 1; i < len(cols); i++ {
			if cols[i-1][0] == cols[i][0] {
				continue
			}
			v := w + cols[i-1][0]
			for j := i - 1; j >= 0; j-- {
				if cols[j][0] != cols[i-1][0] {
					break
				}
				g[cols[j][1]] = append(g[cols[j][1]], v)
			}
			for j := i; j < len(cols); j++ {
				if cols[j][0] != cols[i][0] {
					break
				}
				g[v] = append(g[v], cols[j][1])
			}
		}
	}
	n := w + h*w
	used := make([]bool, n)
	var topo []int
	var dfs func(x int)
	dfs = func(x int) {
		used[x] = true
		for _, next := range g[x] {
			if used[next] {
				continue
			}
			dfs(next)
		}
		topo = append(topo, x)
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			dfs(i)
		}
	}
	for i := 0; i < len(topo)/2; i++ {
		j := len(topo) - 1 - i
		topo[i], topo[j] = topo[j], topo[i]
	}
	for i := 0; i < n; i++ {
		used[i] = false
	}

	for _, v := range topo {
		used[v] = true
		for _, next := range g[v] {
			if used[next] {
				out("No")
				return
			}
		}
	}

	out("Yes")
}
