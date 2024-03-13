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

type list struct {
	prev, next int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	start := a[0]
	m := make(map[int]list)
	for i := 0; i < N; i++ {
		p := list{-1, -1}
		if i != 0 {
			p.prev = a[i-1]
		}
		if i != N-1 {
			p.next = a[i+1]
		}
		m[a[i]] = p
	}

	Q := getI()
	for qi := 0; qi < Q; qi++ {
		t := getI()
		if t == 1 {
			x, y := getI(), getI()
			nx := m[x].next
			p := list{x, nx}
			m[y] = p
			p = list{m[x].prev, y}
			m[x] = p
			if nx != -1 {
				p = list{y, m[nx].next}
				m[nx] = p
			}
			// out("t==1")
			// cur := start
			// for cur != -1 {
			// 	fmt.Fprint(wr, cur, " ")
			// 	cur = m[cur].next
			// }
			// out()
		} else {
			x := getI()
			px := m[x].prev
			nx := m[x].next
			if start == x {
				start = nx
			}
			if px != -1 {
				m[px] = list{m[px].prev, nx}
			}
			if nx != -1 {
				m[nx] = list{px, m[nx].next}
			}
			// out("t==2")
			// cur := start
			// for cur != -1 {
			// 	fmt.Fprint(wr, cur, " ")
			// 	cur = m[cur].next
			// }
			// out()
		}
	}
	cur := start
	for cur != -1 {
		fmt.Fprint(wr, cur, " ")
		cur = m[cur].next
	}
	out()
}
