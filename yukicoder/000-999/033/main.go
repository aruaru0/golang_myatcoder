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
	N, D, T := getInt(), getInt(), getInt()
	x := getInts(N)
	m := make(map[int][]int)
	for i := 0; i < N; i++ {
		pos := (x[i] + 1e9) % D
		m[pos] = append(m[pos], x[i])
	}
	for i := range m {
		sort.Ints(m[i])
	}
	dt := D * T
	ans := 0
	for _, e := range m {
		f := e[0]
		t := e[0]
		cnt := 0
		for i := 1; i < len(e); i++ {
			if t+2*dt >= e[i] {
				t = e[i]
			} else {
				// out(f, t, e[i], dt)
				cnt += 1 + (t-f)/D + T*2
				f = e[i]
				t = e[i]
			}
		}
		cnt += 1 + (t-f)/D + T*2
		ans += cnt
	}
	// out(m)
	out(ans)
}
