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
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	a := make([]int, N)
	b := make([]int, N)
	m := make(map[int]int)
	mm := make(map[int]int)
	for i := 0; i < N; i++ {
		x, y := getI(), getI()
		a[i], b[i] = x, y
		m[x]++
		m[y]++
		mm[x]++
	}
	c := make([]int, 0)
	for e := range m {
		c = append(c, e)
	}
	sort.Ints(c)
	for i, e := range c {
		m[e] = i
	}
	p := make([]int, len(c)+1)
	for i := 0; i < N; i++ {
		l := m[a[i]]
		r := m[b[i]]
		p[l]++
		p[r+1]--
	}
	for i := 1; i <= len(c); i++ {
		p[i] += p[i-1]
	}

	cnt := p[0]
	// out(c)
	// out(p)
	for i := 1; i < len(c); i++ {
		v := (c[i] - c[i-1] - 1)
		cnt += v*(p[i]-mm[c[i]]) + p[i]
		// out(cnt, v, n, "c", c[i-1], c[i])
		if cnt > K {
			diff := cnt - K
			// out(diff)
			ans := c[i]
			if diff >= p[i] {
				ans--
				diff -= p[i]
			}
			vv := p[i] - mm[c[i]]
			// out(diff, ans, vv)
			ans -= diff / vv
			out(ans)
			return
		}
	}
}
