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

const inf int = 1e18

type pair struct {
	x, y   int
	px, py int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	ax, ay := getI()-1, getI()-1
	bx, by := getI()-1, getI()-1
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
	}
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = inf
		}
	}

	dist[ax][ay] = 0
	q := []pair{pair{ax, ay, -1, -1}}
	dx := []int{-1, -1, 1, 1}
	dy := []int{-1, 1, -1, 1}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			cx, cy := cur.x, cur.y
			for {
				px := cx + dx[i]
				py := cy + dy[i]
				// 来た方に戻る方向はなし
				if px == cur.px && py == cur.py {
					break
				}
				// 範囲外はなし
				if px < 0 || px >= N || py < 0 || py >= N {
					break
				}
				// 壁はなし
				if s[px][py] == '#' {
					break
				}
				// 既に訪れて、かつ、手数が大きい場合はなし
				if dist[px][py] != inf && dist[px][py] < dist[cur.x][cur.y]+1 {
					break
				}
				dist[px][py] = dist[cur.x][cur.y] + 1
				q = append(q, pair{px, py, cx, cy})
				cx, cy = px, py
			}
		}
		// for i := 0; i < N; i++ {
		// 	out(dist[i])
		// }
	}

	if dist[bx][by] == inf {
		out(-1)
	} else {
		out(dist[bx][by])
	}
}
