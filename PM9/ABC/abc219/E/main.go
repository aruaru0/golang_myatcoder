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
	x, y int
}

func bsf(x, y, cnt int, used [6][6]bool, p [6][6]int) [6][6]bool {
	col := p[y][x]
	used[y][x] = true
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	q := []pos{pos{x, y}}
	for len(q) != 0 {
		cx, cy := q[0].x, q[0].y
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px < 0 || px >= 6 || py < 0 || py >= 6 {
				continue
			}
			if p[py][px] != col {
				continue
			}
			if used[py][px] == false {
				used[py][px] = true
				q = append(q, pos{px, py})
			}
		}
	}
	return used
}

func check(bit int) bool {
	var p [6][6]int
	// 壁で囲まれた部分をマーク
	for y := 4; y >= 1; y-- {
		for x := 4; x >= 1; x-- {
			if bit&1 == 1 {
				p[y][x] = 1
			}
			bit >>= 1
		}
	}
	// つながっている部分をカウント
	var used [6][6]bool
	cnt := 0
	for y := 0; y < 6; y++ {
		for x := 0; x < 6; x++ {
			if used[y][x] == false {
				used = bsf(x, y, cnt, used, p)
				cnt++
			}
		}
	}
	// 壁内と壁外の２ブロックから構成されていれば条件を満たす
	if cnt == 2 {
		return true
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	a := make([][]int, 4)
	pat := 0
	for i := 0; i < 4; i++ {
		a[i] = getInts(4)
		for j := 0; j < 4; j++ {
			pat = pat<<1 | a[i][j]
		}
	}

	n := 1 << 16
	cnt := 0
	for bit := 0; bit < n; bit++ {
		// すべての家を壁が含まない場合、continue
		if bit&pat != pat {
			continue
		}
		ok := check(bit)
		if ok {
			cnt++
		}
		// fmt.Printf("%x %x\n", bit, pat)
	}

	out(cnt)
}
