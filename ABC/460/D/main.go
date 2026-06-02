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

type P struct {
	r, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([][]byte, H)
	for i := 0; i < H; i++ {
		s[i] = []byte(getS())
	}
	di := []int{-1, 0, 1, 0, -1, -1, 1, 1}
	dj := []int{0, -1, 0, 1, -1, 1, -1, 1}

	for ti := 0; ti < 2; ti++ {
		t := make([][]byte, H)
		for i := 0; i < H; i++ {
			t[i] = make([]byte, W)
			copy(t[i], s[i])
		}

		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				s[i][j] = '.'
				if t[i][j] == '.' {
					for v := 0; v < 8; v++ {
						ni, nj := i+di[v], j+dj[v]
						if ni < 0 || ni >= H || nj < 0 || nj >= W {
							continue
						}
						if t[ni][nj] == '#' {
							s[i][j] = '#'
							break
						}
					}
				}
			}
		}
	}

	const inf = 1e18 + 1
	q := make([]P, 0)
	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			if s[i][j] == '#' {
				q = append(q, P{i, j})
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}
	for len(q) != 0 {
		i, j := q[0].r, q[0].c
		q = q[1:]
		for v := 0; v < 8; v++ {
			ni, nj := i+di[v], j+dj[v]
			if ni < 0 || ni >= H || nj < 0 || nj >= W {
				continue
			}
			if dist[ni][nj] != inf {
				continue
			}
			dist[ni][nj] = dist[i][j] + 1
			q = append(q, P{ni, nj})
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if dist[i][j]%2 == 0 {
				s[i][j] = '#'
			} else {
				s[i][j] = '.'
			}
		}
		out(string(s[i]))
	}

}
