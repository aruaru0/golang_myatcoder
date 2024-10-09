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

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

var N, S, T int
var a, b, c, d []int
var t []float64
var p []int

func calc(bit int) float64 {
	sx, sy := 0, 0
	s := float64(S)
	// fmt.Fprintf(wr, "%3b:\n", bit)
	tot := 0.0
	for _, idx := range p {
		nx, ny := a[idx], b[idx]
		lx, ly := c[idx], d[idx]
		if (bit>>idx)%2 == 1 {
			nx, lx = lx, nx
			ny, ly = ly, ny
		}
		diff := math.Sqrt(math.Pow(float64(nx-sx), 2) + math.Pow(float64(ny-sy), 2))
		tot += diff/s + t[idx]
		// out(sx, sy, "->", nx, ny, "->", lx, ly, "s", diff, tot)
		sx, sy = lx, ly
	}
	return tot
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, S, T = getI(), getI(), getI()
	a, b, c, d = make([]int, N), make([]int, N), make([]int, N), make([]int, N)
	t = make([]float64, N)
	for i := 0; i < N; i++ {
		a[i], b[i], c[i], d[i] = getI(), getI(), getI(), getI()
		dx := float64(a[i] - c[i])
		dy := float64(b[i] - d[i])
		t[i] = math.Sqrt(dx*dx+dy*dy) / float64(T)
	}

	p = make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = i
	}

	ans := math.MaxFloat64
	for {
		// out(p)
		for bit := 0; bit < 1<<N; bit++ {
			ret := calc(bit)
			ans = math.Min(ans, ret)
			// out(ret)
		}

		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
	}
	out(ans)
}
