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

	n := 1 << 3
	ans := 0
	for i := 0; i < n; i++ {
		flag := make([]int, 3)
		for j := 0; j < 3; j++ {
			if (i>>j)%2 == 1 {
				flag[j] = 1
			} else {
				flag[j] = -1
			}
		}
		sort.Slice(c, func(i, j int) bool {
			x := flag[0]*c[i].x + flag[1]*c[i].y + flag[2]*c[i].z
			y := flag[0]*c[j].x + flag[1]*c[j].y + flag[2]*c[j].z
			return x > y
		})
		x, y, z := 0, 0, 0
		for j := 0; j < M; j++ {
			x += flag[0] * c[j].x
			y += flag[1] * c[j].y
			z += flag[2] * c[j].z
		}
		ans = max(ans, x+y+z)
	}
	out(ans)
}
