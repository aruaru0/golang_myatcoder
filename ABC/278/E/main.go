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
	H, W, N, h, w := getI(), getI(), getI(), getI(), getI()
	a := make([][]int, H)
	d := make([][][300]int, H+1)
	for i := 0; i <= H; i++ {
		if i != H {
			a[i] = getInts(W)
		}
		d[i] = make([][300]int, W+1)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			d[i+1][j+1][a[i][j]-1]++
		}
	}
	// for n := 0; n < N; n++ {
	// 	out("----", n+1)
	// 	for i := 0; i <= H; i++ {
	// 		for j := 0; j <= W; j++ {
	// 			fmt.Fprint(wr, d[i][j][n], " ")
	// 		}
	// 		out("")
	// 	}
	// }

	for i := 0; i <= H; i++ {
		for j := 0; j < W; j++ {
			for n := 0; n < N; n++ {
				d[i][j+1][n] += d[i][j][n]
			}
		}
	}
	for j := 0; j <= W; j++ {
		for i := 0; i < H; i++ {
			for n := 0; n < N; n++ {
				d[i+1][j][n] += d[i][j][n]
			}
		}
	}
	// for n := 0; n < N; n++ {
	// 	out("----", n+1)
	// 	for i := 0; i <= H; i++ {

	// 		for j := 0; j <= W; j++ {
	// 			fmt.Fprint(wr, d[i][j][n], " ")
	// 		}
	// 		out("")
	// 	}
	// }

	for i := 0; i <= H-h; i++ {
		for j := 0; j <= W-w; j++ {
			cnt := 0
			for n := 0; n < N; n++ {
				sum := d[i+h][j+w][n] - d[i][j+w][n] - d[i+h][j][n] + d[i][j][n]
				tot := d[H][W][n] - d[H][0][n] - d[0][W][n] + d[0][0][n]
				if tot-sum > 0 {
					cnt++
				}
			}
			fmt.Fprint(wr, cnt, " ")
		}
		out()
	}
}
