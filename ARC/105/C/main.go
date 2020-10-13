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

const inf = int(1e15)

var N, M int
var w []int
var lv [][2]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	w = getInts(N)
	lv = make([][2]int, M)
	for i := 0; i < M; i++ {
		lv[i] = [2]int{getI(), getI()}
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if w[j] > lv[i][1] {
				out(-1)
				return
			}
		}
	}

	sort.Slice(lv, func(i, j int) bool {
		if lv[i][0] == lv[j][0] {
			return lv[i][1] > lv[j][1]
		}
		return lv[i][0] < lv[j][0]
	})

	for i := M - 2; i >= 0; i-- {
		lv[i][1] = min(lv[i][1], lv[i+1][1])
	}

	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}

	ans := math.MaxInt64
	for {
		dp := make([]int, N)
		for i := 0; i < N; i++ {
			ww := w[a[i]]
			for j := i + 1; j < N; j++ {
				ww += w[a[j]]
				idx := sort.Search(M, func(idx int) bool { return lv[idx][1] >= ww }) - 1
				ll := 0
				if idx != -1 {
					ll += lv[idx][0]
				}
				dp[j] = max(dp[j], dp[i]+ll)
			}
		}
		ans = min(ans, dp[N-1])
		if !NextPermutation(sort.IntSlice(a)) {
			break
		}
	}
	out(ans)
}

// NextPermutation generates the next permutation of the
// sortable collection x in lexical order.  It returns false
// if the permutations are exhausted.
//
// Knuth, Donald (2011), "Section 7.2.1.2: Generating All Permutations",
// The Art of Computer Programming, volume 4A.
// ※NextPermutationは辞書順で次を返す
func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
