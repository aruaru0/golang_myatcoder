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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	c := make([][]int, N+1)
	for i := 0; i < N; i++ {
		x, cc := getI(), getI()-1
		c[cc] = append(c[cc], x)
	}
	c[N] = append(c[N], 0)
	for i := 0; i < N; i++ {
		sort.Ints(c[i])
	}
	// out(c)
	pos := []int{0, 0}
	cost := []int{0, 0}
	for _, e := range c {
		if len(e) == 0 {
			continue
		}
		l := e[0]
		r := e[len(e)-1]
		// l -> r
		toR := inf
		for i := 0; i < 2; i++ {
			v := cost[i]
			if l > pos[i] {
				toR = min(toR, v+r-pos[i])
			} else if r < pos[i] {
				toR = min(toR, v+(pos[i]-l)+r-l)
			} else {
				toR = min(toR, v+(pos[i]-l)+r-l)
			}
		}
		// r -> l
		toL := inf
		for i := 0; i < 2; i++ {
			v := cost[i]
			if l > pos[i] {
				toL = min(toL, v+(r-pos[i])+r-l)
			} else if r < pos[i] {
				toL = min(toL, v+pos[i]-l)
			} else {
				toL = min(toL, v+(r-pos[i])+r-l)
			}
		}
		pos[0] = l
		cost[0] = toL
		pos[1] = r
		cost[1] = toR
		// out(e, l, r, pos, cost, "=>R", toR, "=>L", toL)
	}
	out(min(cost[0], cost[1]))
}
