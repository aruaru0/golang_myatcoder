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

type pair struct {
	a, b int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, T, M := getI(), getI(), getI()
	// p := make(map[pair]bool)
	// for i := 0; i < M; i++ {
	// 	a, b := getI()-1, getI()-1
	// 	if a > b {
	// 		a, b = b, a
	// 	}
	// 	p[pair{a, b}] = true
	// }

	// 可能なパターンをpatに保存
	hate := make([]int, N)
	for i := 0; i < M; i++ {
		a, b := getI()-1, getI()-1
		hate[a] |= 1 << b
		hate[b] |= 1 << a
	}

	pat := make([]bool, 1<<N)
	for bit := 0; bit < 1<<N; bit++ {
		ok := true
		for i, h := range hate {
			if bit>>i%2 == 0 {
				continue
			}
			if bit&h != 0 {
				ok = false
			}
		}
		if ok {
			pat[bit] = true
		}
	}

	var dfs func(pat, k int) int

	dfs = func(p1, k int) int {
		if k == T {
			// 全員割り当ててられているならOK
			if p1 == 1<<N-1 {
				return 1
			}
			return 0
		}

		// 割り当てられる一番番号の小さな人を探す
		mi := N + 1
		for i := 0; i < N; i++ {
			if (p1>>i)%2 == 0 {
				chmin(&mi, i)
			}
		}

		c := 0
		for p2 := 1; p2 < 1<<N; p2++ {
			// mi番を割り当てないならNG
			if (p2>>mi)%2 == 0 {
				continue
			}
			// 作成できる組み合わせになければNG
			if !pat[p2] {
				continue
			}
			// すでに割り当てた人が入っていたらNG
			if p1&p2 != 0 {
				continue
			}

			// p1|p2まで確定させて探索
			c += dfs(p1|p2, k+1)
		}

		return c
	}

	out(dfs(0, 0))

}
