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

type S = []int
type MP struct {
	t S
	a []int
}

func NewMP(t S) *MP {
	mp := &MP{t: t}
	mp.a = make([]int, len(mp.t)+1)
	mp.a[0] = -1
	j := -1
	for i := 0; i < len(mp.t); i++ {
		for j != -1 && t[j] != t[i] {
			j = mp.a[j]
		}
		j++
		mp.a[i+1] = j
	}
	return mp
}

// O(t): tの中でsが出現するインデックスをすべて返す
func (mp *MP) FindAll(s S) (res []int) {
	j := 0
	for i := 0; i < len(s); i++ {
		for j != -1 && mp.t[j] != s[i] {
			j = mp.a[j]
		}
		j++
		if j == len(mp.t) {
			res = append(res, i-j+1)
			j = mp.a[j]
		}
	}
	return
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	b := getInts(N)
	c := make([]int, 0)
	d := make([]int, 0)
	a = append(a, a...)
	for i := 1; i < 2*N-1; i++ {
		c = append(c, a[i]^a[i-1])
	}
	for i := 1; i < N; i++ {
		d = append(d, b[i]^b[i-1])
	}

	// out(a, b)
	// out(c, d)

	kmp := NewMP(d)
	x := kmp.FindAll(c)
	for _, e := range x {
		out(e, a[e]^b[0])
	}
}
