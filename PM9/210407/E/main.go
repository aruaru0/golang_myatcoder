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
	A, B, C, D, E, F := getI(), getI(), getI(), getI(), getI(), getI()

	y0 := 0
	y1 := 1
	water := 0
	suger := 0

	for i := 0; A*i <= F/100; i++ {
		for j := 0; B*j <= F/100; j++ {
			if i == 0 && j == 0 {
				continue
			}
			tot := (i*A + j*B) * 100
			if tot > F {
				continue
			}
			rest := F - tot
			// out(i, j, tot)
			for k := 0; k*C <= rest; k++ {
				for l := 0; l*D <= rest; l++ {
					total := tot + C*l + D*k
					if total > F {
						continue
					}
					e0 := 100 * (C*l + D*k)
					e1 := total
					if (100+E)*e0 > 100*E*e1 {
						continue
					}
					// y0/y1 < e0/e1
					if y0*e1 <= e0*y1 {
						y0, y1 = e0, e1
						water = tot
						suger = l*C + k*D
						// out("Update", water, suger)
					}
					// out(e, l, k, total)
				}
			}
		}
	}
	out(water+suger, suger)
}
