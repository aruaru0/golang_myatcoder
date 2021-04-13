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

func score(x []int, y int) int {
	z := make([]int, 10)
	copy(z, x)
	z[y]++
	// out(z)
	tot := 0
	for i := 1; i < 10; i++ {
		x := 1
		for j := 0; j < z[i]; j++ {
			x *= 10
		}
		// out(i, i*x)
		tot += i * x
	}
	return tot
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	k := getI()
	s := getS()
	t := getS()
	a := make([]int, 10)
	b := make([]int, 10)
	c := make([]int, 10)
	for i := 0; i < 10; i++ {
		c[i] = k
	}
	for i := 0; i < 4; i++ {
		x := int(s[i] - '0')
		a[x]++
		c[x]--
	}
	for i := 0; i < 4; i++ {
		x := int(t[i] - '0')
		b[x]++
		c[x]--
	}

	tot := 0
	cnt := 0
	// out(a, b, c)
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i == j && c[i] < 2 {
				continue
			}
			if c[i] < 1 || c[j] < 1 {
				continue
			}
			taka := score(a, i)
			aoki := score(b, j)
			ss := c[i] * c[j]
			if i == j {
				ss = c[i] * (c[i] - 1)
			}
			tot += ss
			if taka > aoki {
				// out(taka, aoki)
				cnt += ss
			}
		}
	}
	out(float64(cnt) / float64(tot))
}
