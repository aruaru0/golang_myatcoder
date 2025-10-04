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
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([][]byte, H)
	used := make([][]bool, H)
	for i := 0; i < H; i++ {
		s[i] = []byte(getS())
		used[i] = make([]bool, W)
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	calc := func(x, y int) int {
		cnt := 0
		for d := 0; d < 4; d++ {
			nx, ny := x+dx[d], y+dy[d]
			if nx < 0 || nx >= W || ny < 0 || ny >= H {
				continue
			}
			if s[ny][nx] == '#' {
				cnt++
			}
		}
		return cnt
	}

	q := make([]pos, 0)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			cnt := calc(j, i)
			if cnt == 2 || s[i][j] == '#' {
				used[i][j] = true
			}
			if cnt == 1 && s[i][j] == '.' {
				q = append(q, pos{j, i})
			}
		}
	}

	// out(q)

	for len(q) != 0 {
		for _, e := range q {
			used[e.y][e.x] = true
			s[e.y][e.x] = '#'
		}
		r := make([]pos, 0)
		for _, e := range q {
			for d := 0; d < 4; d++ {
				nx, ny := e.x+dx[d], e.y+dy[d]
				if nx < 0 || nx >= W || ny < 0 || ny >= H {
					continue
				}
				if s[ny][nx] == '#' {
					continue
				}
				if used[ny][nx] {
					continue
				}
				cnt := calc(nx, ny)
				if cnt == 1 {
					r = append(r, pos{nx, ny})
				}
			}
		}
		q = r
		// out("------")
		// for i := 0; i < H; i++ {
		// 	out(string(s[i]))
		// }
	}

	cnt := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				cnt++
			}
		}
	}
	out(cnt)
}
