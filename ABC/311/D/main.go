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

type pos struct {
	x, y, dir int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
	}

	used := make([][][4]bool, N)
	for i := 0; i < N; i++ {
		used[i] = make([][4]bool, M)
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	q := make([]pos, 0)

	for i := 0; i < 4; i++ {
		used[1][1][i] = true
		q = append(q, pos{1, 1, i})
	}

	for len(q) != 0 {
		cur := q[0]
		cx, cy, dir := cur.x, cur.y, cur.dir
		q = q[1:]

		px := cx + dx[dir]
		py := cy + dy[dir]
		if px < 0 || px >= M || py < 0 || py >= N {
			continue
		}
		if used[py][px][dir] { // 目的地がすでに到達済み
			continue
		}
		if s[py][px] == '#' { // 手前が停止の場合、(cx, cy)から移動していない方向をチェック
			for i := 0; i < 4; i++ {
				if used[cy][cx][i] == false {
					q = append(q, pos{cx, cy, i})
					used[cy][cx][i] = true
				}
			}
			continue
		}
		used[py][px][dir] = true
		q = append(q, pos{px, py, dir})
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			ok := false
			for k := 0; k < 4; k++ {
				if used[i][j][k] {
					ok = true
				}
			}
			if ok {
				ans++
			}
		}
	}

	out(ans)
}
