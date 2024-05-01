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

var H, W int
var s []string
var used [][]int
var cnt [][]int

var x = []int{1, 0, -1, 0}
var y = []int{0, 1, 0, -1}

func dfs(cy, cx, id int) int {
	ret := 1
	used[cy][cx] = id

	if cnt[cy][cx] == 1 {
		return 1
	}

	for i := 0; i < 4; i++ {
		nx := cx + x[i]
		ny := cy + y[i]
		if nx < 0 || nx >= W || ny < 0 || ny >= H {
			continue
		}
		if used[ny][nx] == id {
			continue
		}
		if s[ny][nx] == '#' {
			continue
		}
		ret += dfs(ny, nx, id)
	}

	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}
	cnt = make([][]int, H)
	for i := 0; i < H; i++ {
		cnt[i] = make([]int, W)
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				cnt[i][j] = 1
				continue
			}
			for k := 0; k < 4; k++ {
				nx := j + x[k]
				ny := i + y[k]
				if nx < 0 || nx >= W || ny < 0 || ny >= H {
					continue
				}
				if s[ny][nx] == '#' {
					cnt[i][j] = 1
					break
				}
			}
		}
	}

	used = make([][]int, H)
	for i := 0; i < H; i++ {
		used[i] = make([]int, W)
	}

	ans := 1
	id := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if cnt[i][j] != 0 {
				continue
			}
			if used[i][j] == 0 {
				id++
				ret := dfs(i, j, id)
				// out(i, j, ret)
				ans = max(ans, ret)
			}
		}
	}

	out(ans)
}
