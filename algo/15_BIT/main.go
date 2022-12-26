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

func (b *bit) init(n int) {
	*b = make([]int, n+1)
}

func (b bit) add(i, x int) {
	for idx := i; idx < len(b); idx += idx & -idx {
		b[idx] += x
	}
}

func (b bit) sum(i int) int {
	s := 0
	for idx := i; idx > 0; idx -= idx & -idx {
		s += b[idx]
	}
	return s
}

// a1+a2+a3+...＞＝ｗとなる最小のインデックス
func (b bit) bitlowerBound(w int) int {
	x := 0
	if w <= 0 {
		return 0
	} else {
		r := 1
		for r < len(b) {
			r = r << 1
		}
		for l := r; l > 0; l = l >> 1 {
			if x+l < len(b) && b[x+l] < w {
				w -= b[x+l]
				x += l
			}
		}
	}
	return x + 1
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	var b bit
	b.init(16)
	out(b)
	b.add(1, 1)
	out(b)
	b.add(2, 2)
	b.add(3, 5)
	b.add(4, 10)
	out(b)
	out(b.sum(4))
	out(b.sum(3))
	out(b.bitlowerBound(3))
}
