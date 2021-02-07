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

const inf = int(1e15)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	xf, yf, rf := getF(), getF(), getF()

	X := int(math.Round(xf * 10000))
	Y := int(math.Round(yf * 10000))
	R := int(math.Round(rf * 10000))

	r := R * R

	minX := (X-R)/10000 - 10
	maxX := (X+R)/10000 + 10

	cnt := 0
	for i := minX; i <= maxX; i++ {
		x := i * 10000
		rr := r - (x-X)*(x-X)
		if rr >= 0 {
			rr = int(math.Sqrt(float64(rr)))
			minY := Y - rr
			maxY := Y + rr
			// 誤差修正
			s := minY/10000*10000 - 10000*5
			e := minY/10000*10000 + 10000*5
			ll := inf
			for y := s; y <= e; y += 10000 {
				if (x-X)*(x-X)+(y-Y)*(y-Y) <= r {
					ll = y
					break
				}
			}
			s = maxY/10000*10000 - 10000*5
			e = maxY/10000*10000 + 10000*5
			rr := inf
			for y := e; y >= s; y -= 10000 {
				if (x-X)*(x-X)+(y-Y)*(y-Y) <= r {
					rr = y
					break
				}
			}
			if ll == inf || rr == inf {
				continue
			}
			cnt += (rr-ll)/10000 + 1
		}
	}
	out(cnt)
}
