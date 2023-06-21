package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/emirpasic/gods/trees/redblacktree"
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

func update(tree *redblacktree.Tree, key interface{}, x int) {
	old, found := tree.Get(key)
	if found {
		x += old.(int)
	}
	if x <= 0 {
		tree.Remove(key)
		return
	}
	tree.Put(key, x)
}

func increment(tree *redblacktree.Tree, key interface{}) {
	update(tree, key, 1)
}
func decrement(tree *redblacktree.Tree, key interface{}) {
	update(tree, key, -1)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, k, q := getI(), getI(), getI()

	aa := make([]int, n)
	lt := redblacktree.NewWithIntComparator()
	rt := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		increment(rt, 0)
	}
	for i := 0; i < n-k; i++ {
		increment(lt, 0)
	}

	sum := 0
	for i := 0; i < q; i++ {
		x := getI() - 1
		y := getI()
		a := aa[x]
		_, found := rt.Get(a)
		if found {
			decrement(rt, a)
			sum -= a
			sum += y
			increment(rt, y)
		} else {
			decrement(lt, a)
			increment(lt, y)
		}
		for !rt.Empty() && !lt.Empty() {
			lmax := lt.Right()
			rmin := rt.Left()
			if lmax.Key.(int) <= rmin.Key.(int) {
				break
			}
			decrement(lt, lmax.Key)
			decrement(rt, rmin.Key)
			sum -= rmin.Key.(int)
			sum += lmax.Key.(int)
			increment(lt, rmin.Key)
			increment(rt, lmax.Key)
		}
		aa[x] = y
		out(sum)
	}
}
