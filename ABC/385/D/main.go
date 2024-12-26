package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func outSlice[T any](s []T) {
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

type pair struct {
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	N, M, Sx, Sy := getI(), getI(), getI(), getI()

	X := make(map[int]*rbt.Tree)
	Y := make(map[int]*rbt.Tree)
	for i := 0; i < N; i++ {
		x, y := getI(), getI()
		if X[x] == nil {
			X[x] = rbt.NewWithIntComparator()
		}
		if Y[y] == nil {
			Y[y] = rbt.NewWithIntComparator()
		}
		X[x].Put(y, y)
		Y[y].Put(x, x)
	}

	// out(Sx, Sy)
	cnt := 0
	m := make(map[pair]bool)
	for i := 0; i < M; i++ {
		d, c := getS()[0], getI()
		// out(string(d), c)
		nx, ny := Sx, Sy
		switch d {
		case 'D':
			nx, ny = Sx, Sy-c
			if X[Sx] != nil {
				for {
					l, ok := X[Sx].Ceiling(ny)
					if !ok {
						break
					}
					pos := l.Key.(int)
					if pos > Sy {
						break
					}
					if m[pair{Sx, pos}] != true {
						m[pair{Sx, pos}] = true
						cnt++
					}
					X[Sx].Remove(l.Key)
				}
			}
		case 'U':
			nx, ny = Sx, Sy+c
			if X[Sx] != nil {
				for {
					l, ok := X[Sx].Ceiling(Sy)
					if !ok {
						break
					}
					pos := l.Key.(int)
					if pos > ny {
						break
					}
					if m[pair{Sx, pos}] != true {
						m[pair{Sx, pos}] = true
						cnt++
					}
					X[Sx].Remove(l.Key)
				}
			}
		case 'L':
			nx, ny = Sx-c, Sy
			if Y[Sy] != nil {
				for {
					l, ok := Y[Sy].Ceiling(nx)
					if !ok {
						break
					}
					pos := l.Key.(int)
					if pos > Sx {
						break
					}
					if m[pair{pos, Sy}] != true {
						m[pair{pos, Sy}] = true
						cnt++
					}
					Y[Sy].Remove(l.Key)
				}
			}
		case 'R':
			nx, ny = Sx+c, Sy
			if Y[Sy] != nil {
				for {
					l, ok := Y[Sy].Ceiling(Sx)
					if !ok {
						break
					}
					pos := l.Key.(int)
					if pos > nx {
						break
					}
					if m[pair{pos, Sy}] != true {
						m[pair{pos, Sy}] = true
						cnt++
					}
					Y[Sy].Remove(l.Key)
				}
			}
		}
		Sx, Sy = nx, ny
	}
	out(Sx, Sy, cnt)
}
