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

func small(N, X, M int) {
	a := X
	ans := X
	for i := 2; i <= N; i++ {
		a *= a
		a %= M
		ans += a
	}
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, X, M := getInt(), getInt(), getInt()

	// if N < 1000000000 {
	// 	small(N, X, M)
	// return
	// }

	m := make([]int, M)
	b := make([]int, 0)
	last := -1
	a := X
	b = append(b, a)
	tot := a
	m[a]++
	for i := 1; i < N; i++ {
		a *= a
		a %= M
		if m[a] != 0 {
			last = a
			break
		}
		if a == 0 {
			out(tot)
			return
		}
		tot += a
		m[a]++
		b = append(b, a)
	}
	if last == -1 {
		out(tot)
		return
	}
	start := 0
	ans := 0
	for i := 0; i < len(b); i++ {
		if b[i] == last {
			start = i
			break
		}
		ans += b[i]
		tot -= b[i]
	}
	N -= start
	loop := len(b) - start
	// out(N, loop, tot)
	n := N / loop
	ans += n * tot
	rest := N % loop
	for i := 0; i < rest; i++ {
		ans += b[start+i]
	}
	// out(n*loop + rest)
	out(ans)
}
