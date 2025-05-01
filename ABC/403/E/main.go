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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

type Trie struct {
	to   []map[rune]int
	ans  int
	ng   []bool
	numY []int
}

func NewTrie() *Trie {
	return &Trie{
		to: []map[rune]int{{}}, // 初期ノードを1つ追加
	}
}

func (t *Trie) Add(s string) int {
	v := 0
	for _, c := range s {
		if _, ok := t.to[v][c]; !ok {
			u := len(t.to)
			t.to[v][c] = u
			t.to = append(t.to, map[rune]int{})
		}
		v = t.to[v][c]
	}
	return v
}

func (t *Trie) Init() {
	n := len(t.to)
	t.ng = make([]bool, n)
	t.numY = make([]int, n)
	t.ans = 0
}

func (t *Trie) AddX(v int) {
	if t.ng[v] {
		return
	}
	t.ng[v] = true
	t.ans -= t.numY[v]
	for _, u := range t.to[v] {
		t.AddX(u)
	}
}

func (t *Trie) AddY(v int) {
	if t.ng[v] {
		return
	}
	t.ans++
	t.numY[v]++
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	q := getI()

	t := NewTrie()
	qs := make([][2]int, q)

	for i := 0; i < q; i++ {
		typ := getI()
		s := getS()
		v := t.Add(s)
		qs[i] = [2]int{typ, v}
	}
	t.Init()

	for _, q := range qs {
		if q[0] == 1 {
			t.AddX(q[1])
		} else {
			t.AddY(q[1])
		}
		out(t.ans)
	}
}
