package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

type dist struct {
	l, x, y int
}

func lowerBoundDist(a []dist, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].l >= x
	})
	return idx
}

func upperBoundDist(a []dist, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].l > x
	})
	return idx
}

type unionfind []int

func newUionfind(n int) unionfind {
	ret := make(unionfind, n)
	for i := 0; i < n; i++ {
		ret[i] = -1
	}
	return ret
}

func (u unionfind) r(x int) int {
	if u[x] < 0 {
		return x
	}
	root := u.r(u[x])
	return root
}

func (u unionfind) unite(i, j int) {
	x := u.r(i)
	y := u.r(j)
	if x == y {
		return
	}
	u[x] += u[y]
	u[y] = x
}

func (u unionfind) same(i, j int) bool {
	x := u.r(i)
	y := u.r(j)
	if x == y {
		return true
	}
	return false
}

func (u unionfind) size(i int) int {
	x := u.r(i)
	return -u[x]
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
	}

	d := make([]dist, 0)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			dx := abs(x[i] - x[j])
			dy := abs(y[i] - y[j])
			d = append(d, dist{dx*dx + dy*dy, i, j})
		}
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i].l < d[j].l
	})

	uf := newUionfind(N)

	ans := -1
	for _, e := range d {
		uf.unite(e.x, e.y)
		if uf.same(0, N-1) {
			ans = e.l
			break
		}
	}

	sqr := int(math.Sqrt(float64(ans)))
	if sqr*sqr < ans {
		sqr++
	}

	out((sqr + 9) / 10 * 10)
}
