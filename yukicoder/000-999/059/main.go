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

const maxW = 1000010

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	w := getInts(N)

	var b bit
	b.init(maxW)

	for i := 0; i < N; i++ {
		if w[i] > 0 {
			if b.sum(maxW)-b.sum(w[i]-1) < K {
				b.add(w[i], 1)
			}
		} else {
			if b.sum(-w[i])-b.sum(-w[i]-1) >= 1 {
				b.add(-w[i], -1)
			}
		}
	}
	out(b.sum(maxW))
}
