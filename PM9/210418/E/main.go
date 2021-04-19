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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W, K := getI(), getI(), getI()
	s := make([]string, H)
	for i := 0; i < H; i++ {
		s[i] = getS()
	}

	ans := int(1e10)
	n := 1 << (H - 1)
	for i := 0; i < n; i++ {
		group := make([]int, H)
		g := 0
		for j := 0; j < H-1; j++ {
			group[j] = g
			if (i>>j)%2 == 1 {
				g++
			}
		}
		group[H-1] = g
		// out(group, g)
		tot := g
		g++

		cnt := make([]int, g)
		ng := false
		for j := 0; j < W; j++ {
			tmp := make([]int, g)
			for k := 0; k < H; k++ {
				if s[k][j] == '1' {
					// out(group[k], g)
					tmp[group[k]]++
				}
			}
			ok := true
			for k := 0; k < g; k++ {
				// out(cnt)
				if tmp[k] > K {
					ng = true
				}
				if cnt[k]+tmp[k] > K {
					ok = false
				}
			}
			if !ok {
				// out("----")
				cnt = make([]int, g)
				tot++
			}
			for k := 0; k < g; k++ {
				cnt[k] += tmp[k]
			}
		}
		// out("tot", tot)
		// out(ng)
		if !ng {
			ans = min(ans, tot)
		}
	}
	out(ans)
}
