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
	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt() - 1
	}

	if N >= K {
		pos := 0
		for i := 0; i < K; i++ {
			pos = a[pos]
		}
		out(pos + 1)
		return
	}

	m := make([]int, N)
	pos := 0
	cnt := 0
	for i := 1; i <= N; i++ {
		m[pos] = i
		// out(pos, a[pos])
		pos = a[pos]
		cnt++
		if m[pos] != 0 {
			break
		}
	}
	// out("----")
	// out(m, pos, cnt)
	loop := cnt + 1 - m[pos]
	// out(loop, cnt, m[pos])
	K -= m[pos] - 1
	// out(K)
	K %= loop
	// out("mod", K)
	for i := 0; i < K; i++ {
		pos = a[pos]
	}
	out(pos + 1)
}
