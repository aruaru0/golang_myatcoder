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

var dx = []int{-1, 1, 0, 0, 0}
var dy = []int{0, 0, 0, -1, 1}

func pos(x, y int) int {
	return x + y*4
}

var memo [1 << 16]*float64

func f(S int) float64 {
	if S == 0 {
		return 0
	}
	if memo[S] != nil {
		return *memo[S]
	}
	mi := math.MaxFloat64
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			tot := 0.0
			cnt := 0
			for i := 0; i < 5; i++ {
				px, py := x+dx[i], y+dy[i]
				id := pos(px, py)
				if px >= 0 && px < 4 && py >= 0 && py < 4 && (S>>id)&1 == 1 {
					tot += 1.0 + f(S-(1<<id))
				} else {
					cnt++
				}
			}
			d := float64(cnt)
			v := (tot + d) / (5 - d)
			mi = math.Min(mi, v)
		}
	}
	memo[S] = &mi
	return mi
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	S := 0
	for y := 0; y < 4; y++ {
		s := getS()
		for x := 0; x < 4; x++ {
			if s[x] == '#' {
				S |= 1 << pos(x, y)
			}
		}
	}
	out(f(S))
}
