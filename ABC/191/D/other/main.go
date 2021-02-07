package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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
	// minY := int(Y - R - 1)
	// maxY := int(Y + R + 1)

	// out(minX, maxX, Y, R, r)

	cnt := 0
	for i := minX; i <= maxX; i++ {
		//		for j := minY; j <= maxY; j++ {
		x := i * 10000
		rr := r - (x-X)*(x-X)
		if rr >= 0 {
			rb := big.NewInt(int64(rr))
			rb = rb.Sqrt(rb)
			rr = int(rb.Int64())
			// out("x =", x, "y=", Y-math.Sqrt(rr), Y+math.Sqrt(rr))
			minY := Y - rr
			maxY := Y + rr
			if minY < 0 {
				minY = -abs(minY) / 10000
			} else {
				minY = (minY + 9999) / 10000
			}
			maxY /= 10000
			cnt += maxY - minY + 1
		}
	}
	out(cnt)
}
