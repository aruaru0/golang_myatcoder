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
	N := getI()
	x := make([]float64, N)
	y := make([]float64, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getF(), getF()
	}
	minR := 1e8

	f := func(cx, cy float64) {
		r := 0.0
		for i := 0; i < N; i++ {
			d := math.Hypot(x[i]-cx, y[i]-cy)
			if r < d {
				r = d
			}
		}
		minR = math.Min(minR, r)
	}

	// 2点を直径とする円で条件を満たす最小半径
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			cx := (x[i] + x[j]) / 2
			cy := (y[i] + y[j]) / 2
			f(cx, cy)
		}
	}
	// ３点を円周に持つ円で条件を満たす円の最小半径
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				a := Vec{x[i], y[i]}
				b := Vec{x[j], y[j]}
				c := Vec{x[k], y[k]}
				p := CircumCenter(a, b, c)
				f(p.X, p.Y)
			}
		}
	}

	out(minR)
}

type Vec struct{ X, Y float64 }

func (p Vec) Hypot() float64     { return math.Hypot(p.X, p.Y) }
func (p Vec) Add(q Vec) Vec      { return Vec{p.X + q.X, p.Y + q.Y} }
func (p Vec) Mul(s float64) Vec  { return Vec{p.X * s, p.Y * s} }
func (p Vec) Sub(q Vec) Vec      { return p.Add(q.Mul(-1)) }
func (p Vec) Dist(q Vec) float64 { return p.Sub(q).Hypot() }
func (p Vec) Det(q Vec) float64  { return p.X*q.Y - p.Y*q.X }

// 外接円の座標
func CircumCenter(a, b, c Vec) Vec {
	a1, b1, a2, b2 := b.X-a.X, b.Y-a.Y, c.X-a.X, c.Y-a.Y
	c1, c2, d := a1*a1+b1*b1, a2*a2+b2*b2, 2*(a1*b2-a2*b1)
	return Vec{a.X + (c1*b2-c2*b1)/d, a.Y + (a1*c2-a2*c1)/d}
}

// 外接円の半径
func CircumCenterR(a, b, c Vec) float64 {
	ab, ac := b.Sub(a), c.Sub(a)
	return a.Dist(b) * b.Dist(c) * c.Dist(a) / (2 * math.Abs(ab.Det(ac)))
}
