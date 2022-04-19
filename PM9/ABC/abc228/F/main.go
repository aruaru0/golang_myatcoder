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
	H, W, h1, w1, h2, w2 := getI(), getI(), getI(), getI(), getI(), getI()

	// 白スタンプは黒スタンプに収まるように調整
	// ※はみ出ている部分に意味はない
	h2 = min(h1, h2)
	w2 = min(w1, w2)

	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}
	c := make([][]int, H+1)
	d := make([][]int, H+1)
	// 二次元累積和作成
	for i := 0; i <= H; i++ {
		c[i] = make([]int, W+1)
		d[i] = make([]int, W+1)
		if i == 0 {
			continue
		}
		for j := 0; j < W; j++ {
			c[i][j+1] = c[i][j] + a[i-1][j]
		}
	}
	for i := 0; i <= W; i++ {
		for j := 0; j < H; j++ {
			c[j+1][i] += c[j][i]
		}
	}

	// d = i,jを先頭にした場合の矩形ないのスコア（白スタンプ）
	for i := 0; i < H-h2+1; i++ {
		for j := 0; j < W-w2+1; j++ {
			d[i][j] = c[i+h2][j+w2] - c[i+h2][j] - c[i][j+w2] + c[i][j]
		}
	}

	// ？？？
	for i := 0; i < H-h2+1; i++ {
		for k := w1 - w2 + 1; k > 1; k -= k / 2 {
			for j := 0; j < W-w2-k/2+1; j++ {
				d[i][j] = max(d[i][j], d[i][j+k/2])
			}
		}
	}

	// ？？？
	for j := 0; j < W-w2+1; j++ {
		for k := h1 - h2 + 1; k > 1; k -= k / 2 {
			for i := 0; i < H-h2-k/2+1; i++ {
				d[i][j] = max(d[i][j], d[i+k/2][j])
			}
		}
	}

	// 累積和から矩形スコアを削る
	z := 0
	for i := 0; i < H-h1+1; i++ {
		for j := 0; j < W-w1+1; j++ {
			z = max(z, c[i+h1][j+w1]-c[i][j+w1]-c[i+h1][j]+c[i][j]-d[i][j])
		}
	}

	out(z)
}
