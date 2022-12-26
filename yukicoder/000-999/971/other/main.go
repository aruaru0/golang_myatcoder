package main

import (
	"bufio"
	"fmt"
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

var H, W int
var s []string
var dist [][]int
var dx = []int{1, 0}
var dy = []int{0, 1}

const inf = int(1e15)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	dist := make([][]int, H+1)
	for i := 0; i < H+1; i++ {
		dist[i] = make([]int, W+1)
		for j := 0; j < W+1; j++ {
			dist[i][j] = inf
		}
	}
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			if i == 1 && j == 1 {
				dist[i][j] = 0
				continue
			}
			if s[i-1][j-1] == '.' {
				dist[i][j] = min(dist[i][j-1], dist[i-1][j]) + 1
			} else {
				dist[i][j] = min(dist[i][j-1], dist[i-1][j]) + (j - 1) + (i - 1) + 1
			}
		}
	}
	out(dist[H][W])
}
