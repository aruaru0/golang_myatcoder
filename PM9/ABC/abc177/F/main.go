package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/utils"
)

type S = int

var SComparator = utils.IntComparator

type MultiSet struct{ tree *rbt.Tree }

func NewMultiSet() *MultiSet {
	return &MultiSet{rbt.NewWith(SComparator)}
}

func (ms *MultiSet) Put(x S) {
	if values, found := ms.tree.Get(x); found {
		ms.tree.Put(x, append(values.([]S), x))
	} else {
		ms.tree.Put(x, []S{x})
	}
}

func (ms *MultiSet) Remove(x S) {
	values, found := ms.tree.Get(x)
	if !found {
		return
	}
	xs := values.([]S)
	if len(xs) == 1 {
		ms.tree.Remove(x)
	} else {
		ms.tree.Put(x, xs[1:])
	}
}

func (ms *MultiSet) Left() []S {
	if left := ms.tree.Left(); left != nil {
		return left.Value.([]S)
	}
	return nil
}

func (ms *MultiSet) Right(x S) []S {
	if right := ms.tree.Right(); right != nil {
		return right.Value.([]S)
	}
	return nil
}

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
	H, W := getI(), getI()

	// m[今] = 開始
	m := rbt.NewWithIntComparator()

	// mset[今-開始]
	mset := NewMultiSet()

	for i := 0; i < W; i++ {
		m.Put(i, i)
		mset.Put(0)
	}

	for i := 0; i < H; i++ {
		A, B := getI()-1, getI()
		maxStart := -1
		for {
			node, found := m.Ceiling(A)
			if !found || node.Key.(int) > B {
				break
			}
			maxStart = Max(maxStart, node.Value.(int))
			mset.Remove(node.Key.(int) - node.Value.(int))
			m.Remove(node.Key)
		}
		if maxStart != -1 && B < W {
			m.Put(B, maxStart)
			mset.Put(B - maxStart)
		}

		if values := mset.Left(); len(values) > 0 {
			out(i + 1 + values[0])
		} else {
			out(-1)
		}
	}
}

func Max(xs ...int) int {
	max := xs[0]
	for _, x := range xs[1:] {
		if max < x {
			max = x
		}
	}
	return max
}
