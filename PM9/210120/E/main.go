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

type cake struct {
	x, y, z int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	c := make([]cake, N)
	for i := 0; i < N; i++ {
		c[i] = cake{getI(), getI(), getI()}
	}
	// c = append(c, c...)

	ans := 0
	for i := 0; i < 8; i++ {
		a := []int{}
		for k := 0; k < 3; k++ {
			if (i>>k)&1 == 1 {
				a = append(a, 1)
			} else {
				a = append(a, -1)
			}
		}
		sort.Slice(c, func(i, j int) bool {
			x := c[i].x*a[0] + c[i].y*a[1] + c[i].z*a[2]
			y := c[j].x*a[0] + c[j].y*a[1] + c[j].z*a[2]
			return x > y
		})
		sx, sy, sz := 0, 0, 0
		for j := 0; j < M; j++ {
			sx += c[j].x
			sy += c[j].y
			sz += c[j].z
		}
		// out(c, abs(sx)+abs(sy)+abs(sz))
		ans = max(ans, abs(sx)+abs(sy)+abs(sz))
	}
	out(ans)
}
