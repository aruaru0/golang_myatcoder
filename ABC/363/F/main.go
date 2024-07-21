package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

func toRevInt(s string) int {
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		ret = ret*10 + int(s[i]-'0')
	}
	return ret
}

func isInv(s string) bool {
	t := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			return false
		}
		t += string(s[i])
	}
	if s == t {
		return true
	}
	return false
}

var p [][2]int

func rec(n int) string {
	s := strconv.Itoa(n)
	if isInv(s) {
		return strconv.Itoa(n)
	}

	for _, e := range p {
		x, y := e[0], e[1]
		xy := x * y
		if n%xy != 0 {
			continue
		}
		s := rec(n / xy)
		if s == "-1" {
			continue
		}
		return strconv.Itoa(x) + "*" + s + "*" + strconv.Itoa(y)
	}
	return "-1"
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()

	p = make([][2]int, 0)
	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		s := strconv.Itoa(i)
		if strings.Index(s, "0") != -1 {
			continue
		}
		j := toRevInt(s)
		if n%j != 0 {
			continue
		}
		if i <= j {
			p = append(p, [2]int{i, j})
		}
	}

	out(rec(n))
}
