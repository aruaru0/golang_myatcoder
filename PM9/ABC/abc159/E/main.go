package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

func f(pos []int) int {
	// out("----", pos)
	n := len(pos) - 1
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		f, t := pos[i], pos[i+1]
		m[i] = make([]int, W)
		for w := 0; w < W; w++ {
			cnt := 0
			for h := f; h < t; h++ {
				if s[h][w] == '1' {
					cnt++
				}
			}
			m[i][w] = cnt
		}
	}

	tot := make([]int, n)
	cnt := 0
	for w := 0; w < W; w++ {
		flg := false
		for h := 0; h < n; h++ {
			if tot[h]+m[h][w] > K {
				flg = true
			}
		}
		if flg {
			for h := 0; h < n; h++ {
				tot[h] = m[h][w]
				if tot[h] > K {
					return inf
				}
			}
			cnt++
		} else {
			for h := 0; h < n; h++ {
				tot[h] += m[h][w]
			}
		}
	}
	// for i := 0; i < n; i++ {
	// 	out(m[i])
	// }
	// out(cnt, n-1)
	return cnt + n - 1
}

var H, W, K int
var s []string

const inf = int(1e10)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, K = getI(), getI(), getI()
	s = make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	n := 1 << (H - 1)

	ans := inf
	for h := 0; h < n; h++ {
		pos := []int{0}
		for i := 0; i < H-1; i++ {
			if (h>>i)%2 == 1 {
				pos = append(pos, i+1)
			}
		}
		pos = append(pos, H)
		ans = min(ans, f(pos))
	}
	out(ans)
}
