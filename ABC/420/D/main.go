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

type pos struct {
	x, y, flg int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := getStrings(H)

	dist := make([][][2]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([][2]int, W)
	}

	const inf = int(1e18)
	sx, sy := 0, 0
	gx, gy := 0, 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if s[i][j] == 'S' {
				sx, sy = j, i
			}
			if s[i][j] == 'G' {
				gx, gy = j, i
			}
			dist[i][j] = [2]int{inf, inf}
		}
	}

	// out(sx, sy, "->", gx, gy)

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	q := []pos{{sx, sy, 0}}
	dist[sy][sx][0] = 0
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := cur.x+dx[i], cur.y+dy[i]
			if nx < 0 || nx >= W || ny < 0 || ny >= H {
				continue
			}
			if s[ny][nx] == '#' { //　壁
				continue
			}
			flg := cur.flg
			if s[ny][nx] == 'o' && cur.flg == 1 {
				continue
			}
			if s[ny][nx] == 'x' && cur.flg == 0 {
				continue
			}
			if s[ny][nx] == '?' { // スイッチ
				flg = 1 - flg // スイッチを押した場合は逆面へ
			}
			if dist[ny][nx][flg] > dist[cur.y][cur.x][cur.flg]+1 {
				// out(nx, ny, flg, cur)
				dist[ny][nx][flg] = dist[cur.y][cur.x][cur.flg] + 1
				q = append(q, pos{nx, ny, flg})
			}
		}
	}

	// for f := 0; f < 2; f++ {
	// 	out("-----")
	// 	for i := 0; i < H; i++ {
	// 		for j := 0; j < W; j++ {
	// 			fmt.Fprint(wr, dist[i][j][f], " ")
	// 		}
	// 		out()
	// 	}
	// }

	ans := min(dist[gy][gx][0], dist[gy][gx][1])
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}
