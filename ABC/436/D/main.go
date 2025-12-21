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
	r, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := getStrings(H)
	m := make(map[byte][]pos)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] != '.' && s[i][j] != '#' {
				m[s[i][j]] = append(m[s[i][j]], pos{i, j})
			}
		}
	}

	used := make([][]bool, H)
	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		used[i] = make([]bool, W)
		dist[i] = make([]int, W)
	}

	q := []pos{pos{0, 0}}
	used[0][0] = true

	dr := []int{1, -1, 0, 0}
	dc := []int{0, 0, 1, -1}

	usedWarp := make(map[byte]bool)

	for len(q) != 0 {
		r, c := q[0].r, q[0].c
		q = q[1:]
		if s[r][c] != '.' && usedWarp[s[r][c]] == false {
			usedWarp[s[r][c]] = true
			for _, e := range m[s[r][c]] {
				if used[e.r][e.c] {
					continue
				}
				used[e.r][e.c] = true
				dist[e.r][e.c] = dist[r][c] + 1
				q = append(q, e)
			}
		}

		for i := 0; i < 4; i++ {
			nr, nc := r+dr[i], c+dc[i]
			if nr < 0 || nr >= H || nc < 0 || nc >= W {
				continue
			}
			if s[nr][nc] == '#' {
				continue
			}
			if used[nr][nc] {
				continue
			}
			used[nr][nc] = true
			dist[nr][nc] = dist[r][c] + 1
			q = append(q, pos{nr, nc})
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }

	if used[H-1][W-1] {
		out(dist[H-1][W-1])
	} else {
		out(-1)
	}
}
