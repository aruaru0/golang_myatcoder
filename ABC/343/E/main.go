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

// overlap calculates the overlapping length between two line segments
func overlap(x []int) int {
	minX, maxX := x[0], x[0]
	for i := 0; i < len(x); i++ {
		minX = min(minX, x[i])
		maxX = max(maxX, x[i])
	}
	return max(0, minX+L-maxX)
}

// calc calculates the areas of overlap for 1, 2, and 3 cubes
func calc(a1, b1, c1, a2, b2, c2, a3, b3, c3 int) (int, int, int) {
	x, y, z := overlap([]int{a1, a2, a3}), overlap([]int{b1, b2, b3}), overlap([]int{c1, c2, c3})
	vol3 := x * y * z

	x, y, z = overlap([]int{a1, a2}), overlap([]int{b1, b2}), overlap([]int{c1, c2})
	v12 := x * y * z
	x, y, z = overlap([]int{a1, a3}), overlap([]int{b1, b3}), overlap([]int{c1, c3})
	v13 := x * y * z
	x, y, z = overlap([]int{a2, a3}), overlap([]int{b2, b3}), overlap([]int{c2, c3})
	v23 := x * y * z

	vol2 := v12 + v13 + v23 - vol3*3

	vol1 := 7*7*7*3 - vol2*2 - vol3*3

	return vol1, vol2, vol3
}

const L = 7

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	// out(calc(0, 0, 0, 0, 6, 0, 6, 0, 0))

	v := getInts(3)

	a1, b1, c1 := 0, 0, 0
	for a2 := -L; a2 < L; a2++ {
		for b2 := -L; b2 < L; b2++ {
			for c2 := -L; c2 < L; c2++ {
				for a3 := -L; a3 < L; a3++ {
					for b3 := -L; b3 < L; b3++ {
						for c3 := -L; c3 < L; c3++ {
							v0, v1, v2 := calc(a1, b1, c1, a2, b2, c2, a3, b3, c3)
							// out(a2, b2, c2, a3, b3, c3)
							if v0 == v[0] && v1 == v[1] && v2 == v[2] {
								out("Yes")
								out(a1, b1, c1, a2, b2, c2, a3, b3, c3)
								return
							}
						}
					}
				}
			}
		}
	}
	out("No")
}
