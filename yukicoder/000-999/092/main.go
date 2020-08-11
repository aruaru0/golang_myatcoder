package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

type pair struct {
	t, c int
}

var N, M, K int
var node [][]pair
var d []int
var find []int
var memo [100][1010]bool

func check(p, k int) {
	if k == K {
		find = append(find, (p + 1))
		return
	}
	if memo[p][k] == true {
		return
	}
	memo[p][k] = true
	for _, v := range node[p] {
		if v.c == d[k] {
			check(v.t, k+1)
		}
	}
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M, K = getInt(), getInt(), getInt()
	node = make([][]pair, N)
	for i := 0; i < M; i++ {
		f, t, c := getInt()-1, getInt()-1, getInt()
		// out(f, t, c)
		node[f] = append(node[f], pair{t, c})
		node[t] = append(node[t], pair{f, c})
	}
	d = getInts(K)
	for i := 0; i < N; i++ {
		check(i, 0)
	}
	ans := make([]int, 0)
	sort.Ints(find)
	prev := -1
	for _, e := range find {
		if e != prev {
			ans = append(ans, e)
		}
		prev = e
	}
	out(len(ans))
	for _, e := range ans {
		fmt.Print(e, " ")
	}
	out()
}
