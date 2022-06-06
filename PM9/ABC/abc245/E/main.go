package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
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

type rect struct {
	w, h int
}

// いい加減、set, multisetだと実装が簡単なやつは勘弁してください。
// segtreeで座標圧縮でやればできるけど、難易度が大きくちがいません？
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	a := getInts(N)
	b := getInts(N)
	c := getInts(M)
	d := getInts(M)

	choco := make([]rect, N)
	for i := 0; i < N; i++ {
		choco[i] = rect{a[i], b[i]}
	}
	box := make([]rect, M)
	for i := 0; i < M; i++ {
		box[i] = rect{c[i], d[i]}
	}
	// 幅でソート
	sort.Slice(choco, func(i, j int) bool {
		return choco[i].w > choco[j].w
	})
	sort.Slice(box, func(i, j int) bool {
		return box[i].w > box[j].w
	})
	tree := rbt.NewWithIntComparator()
	ans := "Yes"
	bi := 0
	for i := 0; i < N; i++ {
		c := choco[i]
		// 箱の幅がチョコより大きければrbtreeに追加
		for ; bi < M && c.w <= box[bi].w; bi++ {
			b := box[bi]
			if count, ok := tree.Get(b.h); ok {
				// 既にある場合はカウントを１増加
				tree.Put(b.h, count.(int)+1)
			} else {
				// 新規追加
				tree.Put(b.h, 1)
			}
		}

		// 高さが入る箱で一番小さいやつを利用
		if node, ok := tree.Ceiling(c.h); ok {
			height := node.Key.(int)
			count := node.Value.(int)
			if count == 1 {
				tree.Remove(height)
			} else {
				tree.Put(height, count-1)
			}
			continue
		} else {
			ans = "No"
			break
		}
	}
	out(ans)
}
