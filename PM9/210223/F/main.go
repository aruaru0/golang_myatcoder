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

func calcpos(n, N int) pos {
	return pos{n % N, n / N}
}

type pos struct {
	x, y int
}

const inf = int(1e18)

var N int
var t [][]int
var flg [][]bool
var dx []int = []int{-1, 1, 0, 0}
var dy []int = []int{0, 0, 1, -1}

func dfs(s pos) {
	z := t[s.y][s.x]
	if flg[s.y][s.x] == true {
		z++
	}
	for i := 0; i < 4; i++ {
		px := s.x + dx[i]
		py := s.y + dy[i]
		if px < 0 || px >= N || py < 0 || py >= N {
			continue
		}
		if t[py][px] > z {
			t[py][px] = z
			dfs(pos{px, py})
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	p := getInts(N * N)

	flg = make([][]bool, N)
	t = make([][]int, N)
	for i := 0; i < N; i++ {
		t[i] = make([]int, N)
		flg[i] = make([]bool, N)
		for j := 0; j < N; j++ {
			t[i][j] = nmin(i, N-1-i, j, N-1-j)
			flg[i][j] = true
		}
	}

	ans := 0
	for i := 0; i < N*N; i++ {
		s := calcpos(p[i]-1, N)
		ans += t[s.y][s.x]
		flg[s.y][s.x] = false
		dfs(s)
	}
	out(ans)
}
