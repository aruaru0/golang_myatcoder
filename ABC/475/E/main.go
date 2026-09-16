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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type Node struct {
	to  [2]int
	cnt int
}

func newNode() Node {
	return Node{
		to:  [2]int{-1, -1},
		cnt: 0,
	}
}

var trie []Node

func insert(s []int, delta int) {
	curr := 0
	trie[curr].cnt += delta
	for _, b := range s {
		if trie[curr].to[b] == -1 {
			trie[curr].to[b] = len(trie)
			trie = append(trie, newNode())
		}
		curr = trie[curr].to[b]
		trie[curr].cnt += delta
	}
}

func check(s []int, M int) bool {
	curr, rem_M := 0, M
	for _, b := range s {
		c0 := 0
		next0 := trie[curr].to[0]
		if next0 != -1 {
			c0 = trie[next0].cnt
		}
		if b == 0 {
			if c0 <= rem_M {
				return true
			} else {
				curr = next0
			}
		} else {
			if c0 <= rem_M {
				rem_M -= c0
				curr = trie[curr].to[1]
			} else {
				return false
			}
		}
	}
	return false
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N, M, K := getI(), getI(), getI()
	T := getS()

	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, K)
	}

	trie = make([]Node, 0, 4000000)
	trie = append(trie, newNode())

	for i := 0; i < N; i++ {
		S := getS()
		for j := 0; j < K; j++ {
			if S[j] == T[j] {
				A[i][j] = 0
			} else {
				A[i][j] = 1
			}
		}
		insert(A[i], 1)
	}

	Q := getI()
	for q := 0; q < Q; q++ {
		i := getI() - 1
		j := getI() - 1

		insert(A[i], -1)
		A[i][j] ^= 1
		insert(A[i], 1)

		if check(A[i], M) {
			out("Yes")
		} else {
			out("No")
		}
	}
}
