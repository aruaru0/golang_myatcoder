package main

import (
	"bufio"
	"fmt"
	"image"
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

// UnionFind高速
type UnionFind struct {
	d []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.d = make([]int, N)
	for i := 0; i < N; i++ {
		u.d[i] = -1
	}
	return u
}

func (p *UnionFind) root(x int) int {
	if p.d[x] < 0 {
		return x
	}
	p.d[x] = p.root(p.d[x])
	return p.d[x]
}

func (p *UnionFind) unite(x, y int) bool {
	x = p.root(x)
	y = p.root(y)
	if x == y {
		return false
	}
	if p.d[x] > p.d[y] {
		x, y = y, x
	}
	p.d[x] += p.d[y]
	p.d[y] = x
	return true
}

func (p *UnionFind) same(x, y int) bool {
	return p.root(x) == p.root(y)
}

func (p *UnionFind) size(x int) int {
	return -p.d[p.root(x)]
}

type pair struct {
	x, y int
}

// ConvexHull returns the set of points that define the
// convex hull of p in CCW order starting from the left most.
func (p points) ConvexHull() points {
	// From https://en.wikibooks.org/wiki/Algorithm_Implementation/Geometry/Convex_hull/Monotone_chain
	// with only minor deviations.
	sort.Sort(p)
	var h points

	// Lower hull
	for _, pt := range p {
		for len(h) >= 2 && !ccw(h[len(h)-2], h[len(h)-1], pt) {
			h = h[:len(h)-1]
		}
		h = append(h, pt)
	}

	// Upper hull
	for i, t := len(p)-2, len(h)+1; i >= 0; i-- {
		pt := p[i]
		for len(h) >= t && !ccw(h[len(h)-2], h[len(h)-1], pt) {
			h = h[:len(h)-1]
		}
		h = append(h, pt)
	}

	return h[:len(h)-1]
}

// ccw returns true if the three points make a counter-clockwise turn
func ccw(a, b, c image.Point) bool {
	return ((b.X - a.X) * (c.Y - a.Y)) > ((b.Y - a.Y) * (c.X - a.X))
}

type points []image.Point

func (p points) Len() int      { return len(p) }
func (p points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p points) Less(i, j int) bool {
	if p[i].X == p[j].X {
		return p[i].Y < p[i].Y
	}
	return p[i].X < p[j].X
}

// func main() {
// 	pts := points{
// 		{16, 3}, {12, 17}, {0, 6}, {-4, -6}, {16, 6},
// 		{16, -7}, {16, -3}, {17, -4}, {5, 19}, {19, -8},
// 		{3, 16}, {12, 13}, {3, -4}, {17, 5}, {-3, 15},
// 		{-3, -9}, {0, 11}, {-9, -3}, {-4, -2}, {12, 10},
// 	}
// 	pts = append(pts, image.Point{1, 1})

// 	hull := pts.ConvexHull()
// 	fmt.Println("Convex Hull:", hull)
// }

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	if N == 0 {
		out(1)
		return
	}
	x := make([]int, N)
	y := make([]int, N)
	m := make(map[pair][]int)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
		pos := pair{x[i] / 10, y[i] / 10}
		m[pos] = append(m[pos], i)
	}
	uf := newUnionFind(N)
	for p, e := range m {
		for yy := p.y - 1; yy <= p.y+1; yy++ {
			for xx := p.x - 1; xx <= p.x+1; xx++ {
				v := m[pair{xx, yy}]
				for _, ei := range e {
					for _, vi := range v {
						dx := abs(x[ei] - x[vi])
						dy := abs(y[ei] - y[vi])
						if dx*dx+dy*dy <= 100 {
							uf.unite(ei, vi)
						}
					}
				}
				// out(xx, yy, p, e, v)
			}
		}
	}

	g := make(map[int][]int)
	for i := 0; i < N; i++ {
		r := uf.root(i)
		g[r] = append(g[r], i)
	}

	ans := 0
	for _, v := range g {
		if len(v) <= 1 {
			continue
		}
		pts := make(points, 0)
		tx := x[v[0]]
		ty := y[v[0]]
		flgx := true
		flgy := true
		for _, e := range v {
			pts = append(pts, image.Point{x[e], y[e]})
			if x[e] != tx {
				flgx = false
			}
			if y[e] != ty {
				flgy = false
			}
		}
		flg := flgx || flgy //  同一直線状に並んでいる場合はと凸包検知しない
		hull := pts
		if flg == false {
			hull = pts.ConvexHull()
		}
		for _, i := range hull {
			for _, j := range hull {
				xx := abs(i.X - j.X)
				yy := abs(i.Y - j.Y)
				ans = max(ans, xx*xx+yy*yy)
			}
		}
	}
	out(math.Sqrt(float64(ans)) + 2)
}
