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

const inf = 1001001001001

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	a := make([]int, N)
	s := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		s += a[i]
	}
	d := make([]int, 0)
	for i := 1; i*i <= s; i++ {
		if s%i == 0 {
			d = append(d, i)
			if s/i != i {
				d = append(d, s/i)
			}
		}
	}

	ans := 0
	for _, v := range d {
		m := make([]int, N)
		for i := 0; i < N; i++ {
			m[i] = a[i] % v
		}
		sort.Ints(m)
		x := make([]int, N)
		x[0] = (v - m[0]) % v
		for i := 1; i < N; i++ {
			x[i] = (v - m[i]) % v
			m[i] += m[i-1]
		}
		for i := N - 2; i >= 0; i-- {
			x[i] += x[i+1]
		}
		kk := inf
		for i := 0; i < N-1; i++ {
			if m[i] == x[i+1] {
				kk = m[i]
			}
		}
		if kk <= K {
			ans = max(ans, v)
		}
		// out(v, m, x, kk)
	}
	out(ans)
}
