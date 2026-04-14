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

// (r, c)
type pos struct {
	r, c, d int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()

	s := getStrings(H)

	const inf = int(1e18)
	// （r、c）に方向dから突入したときの最小コスト（d=0,1,2,3でそれぞれ上、下、左、右）
	dist := make([][][4]int, H)
	prev := make([][][4]pos, H)
	for i := 0; i < H; i++ {
		dist[i] = make([][4]int, W)
		prev[i] = make([][4]pos, W)
		for j := 0; j < W; j++ {
			dist[i][j] = [4]int{inf, inf, inf, inf}
		}
	}

	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	sr, sc := 0, 0
	gr, gc := 0, 0
	for r := 0; r < H; r++ {
		for c := 0; c < W; c++ {
			if s[r][c] == 'S' {
				sr, sc = r, c
			}
			if s[r][c] == 'G' {
				gr, gc = r, c
			}

		}
	}
	// スタートへは全ての方向から突入
	q := make([]pos, 0)
	for i := 0; i < 4; i++ {
		dist[sr][sc][i] = 0
		q = append(q, pos{sr, sc, i})
	}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nr, nc := cur.r+dy[i], cur.c+dx[i]
			if nr < 0 || nr >= H || nc < 0 || nc >= W {
				continue
			}
			if s[nr][nc] == '#' {
				continue
			}
			if dist[nr][nc][i] != inf {
				continue
			}
			if s[cur.r][cur.c] == 'o' {
				if cur.d == i {
					dist[nr][nc][i] = min(dist[nr][nc][i], dist[cur.r][cur.c][cur.d]+1)
					prev[nr][nc][i] = pos{cur.r, cur.c, cur.d}
					q = append(q, pos{nr, nc, i})
				}
			} else if s[cur.r][cur.c] == 'x' {
				if cur.d != i {
					dist[nr][nc][i] = min(dist[nr][nc][i], dist[cur.r][cur.c][cur.d]+1)
					prev[nr][nc][i] = pos{cur.r, cur.c, cur.d}
					q = append(q, pos{nr, nc, i})
				}
			} else {
				dist[nr][nc][i] = min(dist[nr][nc][i], dist[cur.r][cur.c][cur.d]+1)
				prev[nr][nc][i] = pos{cur.r, cur.c, cur.d}
				q = append(q, pos{nr, nc, i})
			}
		}
	}

	t := pos{gr, gc, 0}
	for i := 0; i < 4; i++ {
		if dist[gr][gc][i] < dist[gr][gc][t.d] {
			t.d = i
		}
	}

	if dist[t.r][t.c][t.d] == inf {
		out("No")
		return
	}

	out("Yes")
	move := []byte{'U', 'D', 'L', 'R'}
	ans := []byte{}
	for t.r != sr || t.c != sc {
		nr, nc, nd := prev[t.r][t.c][t.d].r, prev[t.r][t.c][t.d].c, prev[t.r][t.c][t.d].d
		ans = append(ans, move[t.d])
		t = pos{nr, nc, nd}
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	out(string(ans))

}
