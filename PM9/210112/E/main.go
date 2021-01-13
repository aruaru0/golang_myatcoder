package main

import (
	"bufio"
	"fmt"
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

//
// バグ修正（copyを忘れていた)
// 問題なければこちらの組み合わせ列挙を使う
//
func generateComb(index []int, s, r int, ch chan []int) {
	if r != 0 {
		if s < 0 {
			return
		}
		generateComb(index, s-1, r, ch)
		index[r-1] = s
		generateComb(index, s-1, r-1, ch)
	} else {
		out := make([]int, len(index))
		copy(out, index)
		ch <- out
	}
	return
}

func foreachComb(n, k int, ch chan []int) {
	index := make([]int, k)
	generateComb(index, n-1, k, ch)
	close(ch)
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
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, C := getI(), getI()
	D := make([][]int, C)
	for i := 0; i < C; i++ {
		D[i] = getInts(C)
	}
	c := make([][]int, N)
	for i := 0; i < N; i++ {
		c[i] = getInts(N)
	}

	col := make([][]int, 3)
	for i := 0; i < 3; i++ {
		col[i] = make([]int, C)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			pos := (i + j) % 3
			col[pos][c[i][j]-1]++
		}
	}

	// goを使う
	comb := make(chan []int, 1)
	go foreachComb(C, 3, comb)

	ans := int(1e18)

	// out(col)

	for ch := range comb {
		a := []int{0, 1, 2}

		for {
			// out(ch[a[0]], ch[a[1]], ch[a[2]])
			tot := 0
			for i := 0; i < 3; i++ {
				v := ch[a[i]]
				for j := 0; j < C; j++ {
					tot += col[i][j] * D[j][v]
					// out(v, j, col[i][j], D[j][v])
				}
			}
			ans = min(ans, tot)

			if !NextPermutation(sort.IntSlice(a)) {
				break
			}
		}
	}
	out(ans)
}
