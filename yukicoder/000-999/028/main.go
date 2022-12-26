package main

import (
	"bufio"
	"fmt"
	"math"
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

const mod = 100000009

// PfsMap : 素因数分解し、マップを作成
func PfsMap(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2] = pfs[2] + 1
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i] = pfs[i] + 1
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = pfs[n] + 1
	}

	return pfs
}

func solve(s, n, k, b int) {
	pfs := PfsMap(b)
	p := make([][36]int, n+1)
	// out(pfs)
	x := s
	y := x
	for i := range pfs {
		for y%i == 0 {
			p[0][i]++
			y /= i
		}
	}
	for i := 1; i <= n; i++ {
		x = 1 + (x*x+x*12345)%mod
		y := x
		for j := range pfs {
			for y%j == 0 {
				p[i][j]++
				y /= j
			}
		}
	}

	ans := math.MaxInt64
	for i, v := range pfs {
		t := make([]int, n+1)
		for j := 0; j < n+1; j++ {
			t[j] = p[j][i]
		}
		sort.Ints(t)
		cnt := 0
		for j := 0; j < k; j++ {
			cnt += t[j]
		}
		ans = min(ans, cnt/v)
	}
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	Q := getInt()
	// Q = 1
	for i := 0; i < Q; i++ {
		s, n, k, b := getInt(), getInt(), getInt(), getInt()
		solve(s, n, k, b)
	}
}
