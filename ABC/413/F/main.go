package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type nums interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

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
func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// min for n entry
func nmin[T Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax[T Ordered](a ...T) T {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin[T Ordered](a *T, b T) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax[T Ordered](a *T, b T) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
}

func asub[T nums](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}

func abs[T nums](a T) T {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound[T nums](a []T, x T) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound[T nums](a []T, x T) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

const inf = 1001001001

type pair struct{ r, c int }

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1024*1024)

	h, w, k := getI(), getI(), getI()

	dist := make([][]int, h)
	checked := make([][]bool, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		checked[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			dist[i][j] = inf
		}
	}

	q := make([]pair, 0, k)
	for i := 0; i < k; i++ {
		r, c := getI()-1, getI()-1
		dist[r][c] = 0
		q = append(q, pair{r, c})
	}

	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]

		for v := 0; v < 4; v++ {
			ni, nj := cur.r+di[v], cur.c+dj[v]
			if ni < 0 || nj < 0 || ni >= h || nj >= w {
				continue
			}
			if dist[ni][nj] != inf {
				continue
			}
			if checked[ni][nj] {
				dist[ni][nj] = dist[cur.r][cur.c] + 1
				q = append(q, pair{ni, nj})
			} else {
				checked[ni][nj] = true
			}
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if dist[i][j] != inf {
				ans += dist[i][j]
			}
		}
	}
	out(ans)
}
