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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	var a [110][110][110]int
	for i := 0; i < N; i++ {
		x1, y1, z1, x2, y2, z2 := getI(), getI(), getI(), getI(), getI(), getI()
		for x := x1; x < x2; x++ {
			for y := y1; y < y2; y++ {
				for z := z1; z < z2; z++ {
					a[x][y][z] = i + 1
				}
			}
		}
	}
	ans := make([]map[int]bool, N+1)
	for i := 0; i <= N; i++ {
		ans[i] = make(map[int]bool)
	}
	add := func(i, j int) {
		// i->jの接触と、j->iの接触の両方を登録
		ans[i][j] = true
		ans[j][i] = true
	}
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			for z := 0; z < 100; z++ {
				// 物体がなければ、パス
				if a[x][y][z] == 0 {
					continue
				}
				// ｘ方向に隣が別の物体なら登録
				if a[x+1][y][z] != 0 && a[x+1][y][z] != a[x][y][z] {
					add(a[x][y][z], a[x+1][y][z])
				}
				// y方向に隣が別の物体なら登録
				if a[x][y+1][z] != 0 && a[x][y+1][z] != a[x][y][z] {
					add(a[x][y][z], a[x][y+1][z])
				}
				// z方向に隣が別の物体なら登録
				if a[x][y][z+1] != 0 && a[x][y][z+1] != a[x][y][z] {
					add(a[x][y][z], a[x][y][z+1])
				}
			}
		}
	}

	for i := 1; i <= N; i++ {
		out(len(ans[i]))
	}
}
