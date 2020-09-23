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

func main() {
	sc.Split(bufio.ScanWords)
	N, M, K := getInt(), getInt(), getInt()
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInts(M)
	}

	n0 := (N + 1) / 2
	n1 := N / 2
	// out(n0, n1)
	// part 1
	x := 1
	for i := 0; i < n0; i++ {
		x *= M
	}
	b := make([]int, 0)
	for i := 0; i < x; i++ {
		t := i
		sum := 0
		for j := 0; j < n0; j++ {
			sel := t % M
			sum += a[j][sel]
			t /= M
		}
		b = append(b, sum)
	}
	sort.Ints(b)

	x = 1
	for i := 0; i < n1; i++ {
		x *= M
	}
	c := make([]int, 0)
	for i := 0; i < x; i++ {
		t := i
		sum := 0
		for j := 0; j < n1; j++ {
			sel := t % M
			sum += a[j+n0][sel]
			t /= M
		}
		c = append(c, sum)
	}
	sort.Ints(c)

	// out(b)
	// out(c)
	ans := -1
	for i := 0; i < len(b); i++ {
		v := K - b[i]
		p := lowerBound(c, v)
		// out(b[i], p)
		if p == len(c) {
			p--
		}
		if b[i]+c[p] <= K {
			ans = max(ans, b[i]+c[p])
		}
		if p != 0 && b[i]+c[p-1] <= K {
			ans = max(ans, b[i]+c[p-1])
		}
	}
	if ans == -1 {
		out(-1)
		return
	}
	out(K - ans)
}
