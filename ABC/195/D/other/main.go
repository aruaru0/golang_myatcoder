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

var used [][]bool
var H, W, A, B int

func xy2idx(x, y int) int {
	return y*W + x
}

func idx2xy(idx int) (int, int) {
	return idx % W, idx / W
}

func rec(n, a, b int) int {
	x, y := idx2xy(n)
	if a < 0 || b < 0 {
		return 0
	}
	if n == H*W {
		return 1
	}
	var ret int
	if used[y][x] { // すでに置かれている
		ret += rec(n+1, a, b)
	} else {
		// １枚を選択
		used[y][x] = true
		ret += rec(n+1, a, b-1)
		used[y][x] = false
		// 横に２枚
		used[y][x] = true
		if x+1 < W && used[y][x+1] == false {
			used[y][x+1] = true
			ret += rec(n+1, a-1, b)
			used[y][x+1] = false
		}
		// 縦に２枚
		if y+1 < H && used[y+1][x] == false {
			used[y+1][x] = true
			ret += rec(n+1, a-1, b)
			used[y+1][x] = false
		}
		used[y][x] = false
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, A, B = getI(), getI(), getI(), getI()

	used = make([][]bool, H)
	for i := 0; i < H; i++ {
		used[i] = make([]bool, W)
	}

	out(rec(0, A, B))
}
