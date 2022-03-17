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

type pos struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	S := getS()
	K := getI()

	x, y := 0, 0
	st := make(map[pos]bool)
	st[pos{x, y}] = true
	for _, e := range S {
		switch e {
		case 'U':
			y--
		case 'D':
			y++
		case 'L':
			x--
		case 'R':
			x++
		}
		st[pos{x, y}] = true
	}

	if x == 0 && y == 0 {
		out(len(st))
		return
	}

	ps := make([]pos, 0)
	for p := range st {
		ps = append(ps, p)
	}
	if x == 0 {
		x, y = y, x
		for i := 0; i < len(ps); i++ {
			ps[i].x, ps[i].y = ps[i].y, ps[i].x
		}
	}
	if x < 0 {
		x = -x
		for i := 0; i < len(ps); i++ {
			ps[i].x = -ps[i].x
		}
	}
	if y < 0 {
		y = -y
		for i := 0; i < len(ps); i++ {
			ps[i].y = -ps[i].y
		}
	}

	mp := make(map[pos][]pos)
	for _, p := range ps {
		nx, ny := p.x, p.y
		if nx > 0 {
			i := nx / x
			nx -= i * x
			ny -= y * i
		} else {
			i := (-nx + x - 1) / x
			nx += i * x
			ny += y * i
		}
		mp[pos{nx, ny}] = append(mp[pos{nx, ny}], p)
	}

	ans := 0
	for _, np := range mp {
		d := make([]int, 0)
		for _, p := range np {
			d = append(d, (p.x+len(S))/x)
		}
		sort.Ints(d)
		for i := 0; i < len(d); i++ {
			w := K
			if i+1 < len(d) {
				w = min(w, d[i+1]-d[i])
			}
			ans += w
		}
	}
	out(ans)
}
