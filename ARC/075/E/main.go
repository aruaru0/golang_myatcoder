package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	// fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
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

type bit []int

func (p *bit) init(n int) {
	*p = make([]int, n+1)
}
func (p bit) sum(i int) int {
	if i == 0 {
		return 0
	}
	return p[i] + p.sum(i-(i&-i))
}
func (p bit) add(i, x int) {
	if i >= len(p) {
		return
	}
	p[i] += x
	p.add(i+(i&-i), x)
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt() - K
	}

	s := make([]int, N+1)
	for i := 0; i < N; i++ {
		s[i+1] = s[i] + a[i]
	}

	b := make([]int, N)
	for i := 1; i <= N; i++ {
		b[i-1] = s[i]
	}
	out(s)
	sort.Ints(b)
	out(b)
	bm := make(map[int]int)
	for i := 0; i < N; i++ {
		bm[b[i]] = i + 1
		out(bm)
	}
	var tree bit
	tree.init(N + 2)
	for i := 0; i < N; i++ {
		tree.add(bm[b[i]], 1)
	}
	// tree.add(bm[0], -1)
	out("---")
	out(tree)

	ans := 0
	for i := 1; i <= N; i++ {
		pos := lowerBound(b, s[i-1])
		ans += tree.sum(N+1) - tree.sum(pos)
		tree.add(bm[s[i]], -1)
	}
	fmt.Println(ans)
}
