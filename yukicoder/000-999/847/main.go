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

func divisor(n int) []int {
	ret := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if n/i != i {
				ret = append(ret, n/i)
			}
		}
	}
	return ret
}

func psf(n int) map[int]int {
	ret := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			ret[i]++
			n /= i
		}
	}
	if n != 1 {
		ret[n]++
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K, M := getInt(), getInt(), getInt()

	if N == 1 {
		out(1)
		return
	}

	x := psf(N)
	for i, e := range x {
		x[i] = e * K
	}

	m := make(map[int]int)
	m[1] = 1
	for i, e := range x {
		v := 1
		l := make(map[int]int)
		for j := 0; j < e; j++ {
			v *= i
			if v > M {
				break
			}
			l[v]++
		}
		// out(l, m)
		o := make(map[int]int)
		for s := range m {
			for t := range l {
				if s*t <= M {
					o[s*t]++
				}
			}
		}
		for s := range l {
			m[s]++
		}
		for s := range o {
			m[s]++
		}
	}
	out(len(m))
}
