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
	n, l, r := getI(), getI(), getI()
	n2 := 1 << n
	r++
	const inf = int(1e10) + 1
	dist := make([]int, n2+1)
	pre := make([]int, n2+1)
	for i := 0; i < n2+1; i++ {
		dist[i] = inf
		pre[i] = -1
	}
	q := make([]int, 0)
	q = append(q, l)
	dist[l] = 0
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		push := func(to int) {
			if to < 0 || to > n2 {
				return
			}
			if dist[to] != inf {
				return
			}
			dist[to] = dist[v] + 1
			pre[to] = v
			q = append(q, to)
		}
		for i := 0; i < n+1; i++ {
			push(v - (1 << i))
			push(v + (1 << i))
			if ((v >> i) & 1) == 1 {
				break
			}
		}
	}

	ans := 0
	query := func(s, t int) {
		sign := 1
		if s > t {
			s, t = t, s
			sign = -1
		}
		{
			i, j, w := 0, s, t-s
			for w%2 == 0 {
				j >>= 1
				w >>= 1
				i++
			}
			fmt.Println("? ", i, " ", j)
		}
		x := getI()
		ans = (ans + x*sign + 100) % 100
	}
	for r != l {
		query(pre[r], r)
		r = pre[r]
	}
	fmt.Println("! ", ans)

}
