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

// 拡張GCD
func extGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	d, y, x := extGCD(b, a%b)
	y -= a / b * x
	return d, x, y
}

var x, y, z int

var memo map[int]bool

func rec(n, v int, s []byte) {
	if memo[v] == true {
		return
	}
	memo[v] = true
	if n > 10000 {
		return
	}
	if v > z {
		return
	}
	if v == z {
		out(string(s))
		wr.Flush()
		os.Exit(0)
	}
	rec(n+1, v+x, append(s, "cC"...))
	rec(n+1, v-x, append(s, "cW"...))
	rec(n+1, v+y, append(s, "wC"...))
	rec(n+1, v-y, append(s, "wW"...))
}

const inf = int(11000)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	x, y, z = getI(), getI(), getI()

	if x == 0 && z == 0 {
		out("ccC")
		return
	}
	if y == 0 && z == 0 {
		out("wwC")
		return
	}
	if x == 0 && y == 0 && z != 0 {
		out("NO")
		return
	}
	if z == 0 {
		out("ccW")
		return
	}

	A, B := 0, 0
	mi := inf

	if y == 0 {
		if z%x != 0 {
			out("NO")
			return
		}
		A, B = z/x, 0
		mi = A
	} else if x == 0 {
		if z%y != 0 {
			out("NO")
			return
		}
		A, B = 0, z/y
		mi = B
	} else {
		for a := -5000; a <= 5000; a++ {
			if (z-a*x)%y != 0 {
				continue
			}
			b := (z - a*x) / y
			if mi > abs(a)+abs(b) {
				mi = abs(a) + abs(b)
				A, B = a, b
			}
		}
	}
	if mi >= inf {
		out("NO")
		return
	}

	as := make([]byte, 0)
	for i := 0; i < abs(A); i++ {
		as = append(as, 'c')
	}
	ae := make([]byte, 0)
	for i := 0; i < abs(A)-1; i++ {
		if A < 0 {
			ae = append(ae, 'W')
		} else {
			ae = append(ae, 'C')
		}
	}

	bs := make([]byte, 0)
	for i := 0; i < abs(B); i++ {
		bs = append(bs, 'w')
	}
	be := make([]byte, 0)
	for i := 0; i < abs(B)-1; i++ {
		if B < 0 {
			be = append(be, 'W')
		} else {
			be = append(be, 'C')
		}
	}

	if A == 0 {
		as = append(bs, be...)
	} else if B == 0 {
		as = append(as, ae...)
	} else if A > 0 {
		as = append(bs, as...)
		as = append(as, ae...)
		as = append(as, be...)
		if B < 0 {
			as = append(as, 'W')
		} else {
			as = append(as, 'C')
		}
	} else {
		as = append(as, bs...)
		as = append(as, be...)
		as = append(as, ae...)
		as = append(as, 'W')
	}
	if len(as) > 10000 {
		out("NO")
		return
	}
	// out(x, y)
	out(string(as))
}
