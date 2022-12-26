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
	N, Q := getI(), getI()
	a := getInts(N)
	s := getS()

	var d [32][2]int
	for i := 0; i < 32; i++ {
		for x := 0; x < 2; x++ {
			cur := x
			prev := x
			for j := 0; j < N; j++ {
				if s[j] == '0' {
					cur &= (a[j] >> i) & 1
				} else {
					cur |= (a[j] >> i) & 1
				}
				if cur != prev {
					d[i][x]++
				}
				prev = cur
			}
		}
	}

	t := getInts(Q)
	for i := 0; i < Q; i++ {
		tot := 0
		for j := 0; j < 32; j++ {
			x := (t[i] >> j) & 1
			tot += d[j][x] << j
		}
		out(tot)
	}

	// for i := 0; i < 32; i++ {
	// 	out(d[i])
	// }
	// out(Q)
	// for i := 0; i < N; i++ {
	// 	fmt.Printf("%8.8b\n", a[i])
	// }
}

func f(t int, a []int, s string) {
	b := make([]int, 0)
	b = append(b, t)
	for i := 0; i < len(a); i++ {
		b = append(b, a[i])
	}

	x := b[0]
	ans := 0
	prev := b[0]
	for i := 1; i <= len(s); i++ {
		if s[i-1] == '0' {
			x &= b[i]
		} else {
			x |= b[i]
		}
		ans += abs(prev - x)
		prev = x
	}
	out(ans)
}
