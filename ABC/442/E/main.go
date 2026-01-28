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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type Point struct {
	x, y int
}

func getQuadrant(p Point) int {
	if p.x > 0 && p.y >= 0 {
		return 1
	}
	if p.x <= 0 && p.y > 0 {
		return 2
	}
	if p.x < 0 && p.y <= 0 {
		return 3
	}
	return 4
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()

	p := make([]Point, 0)
	m := make(map[Point][]int)
	for i := 0; i < N; i++ {
		x, y := getI(), getI()
		if x == 0 && y != 0 {
			y /= abs(y)
		} else if y == 0 && x != 0 {
			x /= abs(x)
		} else {
			g := gcd(abs(x), abs(y))
			x /= g
			y /= g
		}
		if len(m[Point{x, y}]) == 0 {
			p = append(p, Point{x, y})
		}
		m[Point{x, y}] = append(m[Point{x, y}], i)
	}

	sort.Slice(p, func(i, j int) bool {
		p1 := p[i]
		p2 := p[j]

		q1 := getQuadrant(p1)
		q2 := getQuadrant(p2)

		if q1 != q2 {
			return q1 > q2
		}

		crossProduct := p1.x*p2.y - p2.x*p1.y

		if crossProduct != 0 {
			return crossProduct < 0
		}

		dist1Sq := p1.x*p1.x + p1.y*p1.y
		dist2Sq := p2.x*p2.x + p2.y*p2.y
		return dist1Sq > dist2Sq
	})

	rev := make(map[int]int)
	a := make([]int, len(p)+1)
	for i := 0; i < len(p); i++ {
		a[i+1] = len(m[p[i]])
		for _, e := range m[p[i]] {
			rev[e] = i
		}
	}

	a = append(a, a[1:]...)
	for i := 0; i < len(a)-1; i++ {
		a[i+1] += a[i]
	}

	for qi := 0; qi < Q; qi++ {
		x, y := getI()-1, getI()-1
		l := rev[x]
		r := rev[y]
		if l > r {
			r += len(p)
		}
		d := a[r+1] - a[l]
		// out(l, r, a, d)
		out(d)

	}
}
