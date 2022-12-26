package main

import (
	"bufio"
	"fmt"
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

func calc(a, b, x, y []int, n int) {
	a[0] = 1
	b[0] = 1
	for i := 0; i < n; i++ {
		b[i+1] = (b[i]*y[i]%mod + 1) % mod
	}
	for i := 0; i < n; i++ {
		v := b[i] * b[i] % mod * x[i] % mod
		a[i+1] = (v + a[i]) % mod
	}
}

func calcy(a, b, x, y []int, n int, s int) {
	for i := s; i < n; i++ {
		e := (b[i]*y[i]%mod + 1) % mod
		v := b[i] * b[i] % mod * x[i] % mod
		v = (v + a[i]) % mod
		if e == b[i+1] && v == a[i+1] {
			break
		}
		b[i+1] = e
		a[i+1] = v
	}
}

func calcx(a, b, x, y []int, n int, s int) {
	v := b[s] * b[s] % mod * x[s] % mod
	a[s+1] = (v + a[s]) % mod
	for i := s + 1; i < n; i++ {
		v := b[i] * b[i] % mod * x[i] % mod
		v = (v + a[i]) % mod
		if a[i+1] == v {
			break
		}
		a[i+1] = v
	}
}

const mod = int(1e9 + 7)

type op struct {
	s    string
	i, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	q := getI()

	a := make([]int, n+1)
	b := make([]int, n+1)
	x := make([]int, n+1)
	y := make([]int, n+1)

	calc(a, b, x, y, n)

	Q := make([]op, q)
	cnt := 0
	for i := 0; i < q; i++ {
		s := getS()
		switch s {
		case "x":
			Q[i] = op{s, getI(), getI()}
		case "y":
			Q[i] = op{s, getI(), getI()}
		case "a":
			Q[i] = op{s, getI(), 0}
			cnt++
		}
	}

	if cnt < 100 {
		for k := 0; k < q; k++ {
			s := Q[k].s
			switch s {
			case "x":
				i, v := Q[k].i, Q[k].v
				x[i] = v
			case "y":
				i, v := Q[k].i, Q[k].v
				y[i] = v
			case "a":
				i := Q[k].i
				calc(a, b, x, y, n)
				out(a[i])
			}
		}
		return
	}

	for k := 0; k < q; k++ {
		s := Q[k].s
		switch s {
		case "x":
			i, v := Q[k].i, Q[k].v
			x[i] = v
			calcx(a, b, x, y, n, i)
			// out("x------")
			// out(a, x)
			// out(b, y)
		case "y":
			i, v := Q[k].i, Q[k].v
			y[i] = v
			calcy(a, b, x, y, n, i)
			// out("y------")
			// out(a, x)
			// out(b, y)
		case "a":
			i := Q[k].i
			out(a[i])
		}
	}
}
