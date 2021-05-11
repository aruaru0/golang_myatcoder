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

	// horizontal
	h, w := H/3, W
	d0 := (h+1)*w - h*w
	if H%3 == 0 {
		d0 = 0
	}
	// vertical
	h, w = H, W/3
	d1 := h*(w+1) - h*w
	if W%3 == 0 {
		d1 = 0
	}

	// out(d0, d1)
	// horoz + vert
	l := 0
	r := H + 1
	for l+1 != r {
		m := (l + r) / 2
		if m*W*2 > (H-m)*W {
			r = m
		} else {
			l = m
		}
	}
	x := l * W
	y := (H - l) * (W / 2)
	z := (H-l)*W - y
	d2 := nmax(abs(x-y), abs(x-z), abs(y-z))
	// out(x, y, z, d2)
	x = r * W
	y = (H - r) * (W / 2)
	z = (H-r)*W - y
	d3 := nmax(abs(x-y), abs(x-z), abs(y-z))
	// out(x, y, z, d3)

	// horoz + vert
	l = 0
	r = W + 1
	for l+1 != r {
		m := (l + r) / 2
		if m*H*2 > (W-m)*H {
			r = m
		} else {
			l = m
		}
	}
	x = l * H
	y = (W - l) * (H / 2)
	z = (W-l)*H - y
	d4 := nmax(abs(x-y), abs(x-z), abs(y-z))
	// out(x, y, z, d4)
	x = r * H
	y = (W - r) * (H / 2)
	z = (W-r)*H - y
	d5 := nmax(abs(x-y), abs(x-z), abs(y-z))
	// out(x, y, z, d5)

	out(nmin(d0, d1, d2, d3, d4, d5))

}
