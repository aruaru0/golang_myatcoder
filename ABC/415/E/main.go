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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}
	p := getInts(H + W - 1)

	const inf = int(1e18)

	f := func(x int) bool {
		// gは各マスの残り
		g := make([][]int, H)
		ok := make([][]bool, H)
		for i := 0; i < H; i++ {
			g[i] = make([]int, W)
			ok[i] = make([]bool, W)
			for j := 0; j < W; j++ {
				g[i][j] = inf
			}
		}

		// 最初から無理なら無理
		x += a[0][0]
		if x < p[0] {
			return false
		}

		g[0][0] = x - p[0]

		ok[0][0] = true

		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if i == 0 && j == 0 {
					continue
				}
				// curは初期値-inf
				cur := -inf
				// 上から来れる場合
				if i != 0 && ok[i-1][j] == true {
					cur = max(cur, g[i-1][j])
				}
				// 左から来れる場合
				if j != 0 && ok[i][j-1] == true {
					cur = max(cur, g[i][j-1])
				}
				// マスの値を足して、pを引く
				cur += a[i][j]
				cur -= p[i+j]
				// もし残りが0以上ならこのマスはOK
				if cur >= 0 {
					ok[i][j] = true
					g[i][j] = min(g[i][j], cur)
				}

			}
		}
		return ok[H-1][W-1]
	}

	l, r := -1, 0
	// rの初期値は、Pの合計値＋１にしておく
	for _, e := range p {
		r += e
	}
	r++

	// 二分探索
	for l+1 != r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	out(r)
}
