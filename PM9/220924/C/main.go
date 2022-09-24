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
	H, W := getI(), getI()
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		a[i] = getInts(W)
	}
	s := make([][]int, H+1)
	for i := 0; i <= H; i++ {
		s[i] = make([]int, W+1)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			s[i+1][j+1] = s[i][j+1] + a[i][j]
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			s[i+1][j+1] += s[i+1][j]
		}
	}

	calc := func(x1, y1, x2, y2 int) int {
		return s[x2+1][y2+1] + s[x1][y1] - s[x1][y2+1] - s[x2+1][y1]
	}

	cnt := 0

	// 全探索
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			ss := calc(0, 0, i, j)
			h := []int{-1}
			w := []int{-1}
			//　wを固定して分割した時に同じサイズにできる数
			for k := 0; k < H; k++ {
				if calc(h[len(h)-1]+1, 0, k, j) == ss {
					h = append(h, k)
				}
			}
			// hを固定して考えて同じサイズにできる数
			for k := 0; k < W; k++ {
				if calc(0, w[len(w)-1]+1, i, k) == ss {
					w = append(w, k)
				}
			}
			// 割り切れていないなら分割できないとする
			if h[len(h)-1] != H-1 || w[len(w)-1] != W-1 {
				continue
			}
			// このりの区画を含め同じになっているか確認
			f := true
			for k := 0; k < len(h)-1; k++ {
				for l := 0; l < len(w)-1; l++ {
					if calc(h[k]+1, w[l]+1, h[k+1], w[l+1]) != ss {
						f = false
					}
				}
			}
			if f {
				cnt++
			}
		}
	}
	out(cnt - 1)
}
