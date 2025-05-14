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

type pair struct {
	r, c int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	const inf = int(1e7)
	dist := make([][]int, H)
	dirc := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		dirc[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = inf
			dirc[i][j] = -1
		}
	}

	q := make([]pair, 0)
	for r := 0; r < H; r++ {
		for c := 0; c < W; c++ {
			if s[r][c] == 'E' {
				q = append(q, pair{r, c})
				dist[r][c] = 0
			}
		}
	}
	//          0  1  2  3
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}
	arrow := "v^><"

	for len(q) != 0 {
		cr, cc := q[0].r, q[0].c
		q = q[1:]
		for i := 0; i < 4; i++ {
			pr, pc := cr+dr[i], cc+dc[i]
			if pr < 0 || pr >= H || pc < 0 || pc >= W {
				continue
			}
			if s[pr][pc] == '#' {
				continue
			}
			if dist[pr][pc] > dist[cr][cc]+1 {
				dist[pr][pc] = dist[cr][cc] + 1
				dirc[pr][pc] = i
				q = append(q, pair{pr, pc})
			}
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(dist[i])
	// }
	// out("----")
	// for i := 0; i < H; i++ {
	// 	out(dirc[i])
	// }

	for i := 0; i < H; i++ {
		t := make([]byte, W)
		for j := 0; j < W; j++ {
			if s[i][j] == '.' {
				t[j] = arrow[dirc[i][j]]
			} else {
				t[j] = s[i][j]
			}
		}
		out(string(t))
	}
}
