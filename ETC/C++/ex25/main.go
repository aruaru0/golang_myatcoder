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

// AとBに共通して含まれる要素からなる集合を返す
func intersection(A, B int) int {
	return A & B
}

// AとBのうち少なくとも一方に含まれる要素からなる集合を返す
func unionSet(A, B int) int {
	return A | B
}

// AとBのうちどちらか一方にだけ含まれる要素からなる集合を返す
func symmetricDiff(A, B int) int {
	return A ^ B
}

// Aから値xを除く
func subtract(A int, x int) int {
	return A ^ (1 << x)
}

// Aに含まれる要素に1を加える(ただし、要素49が含まれる場合は0になるものとする)
func increment(A int) int {
	last := (A >> 49) & 1
	A = A<<1 + last

	return A
}

// Aに含まれる要素から1を引く(ただし、要素0が含まれる場合は49になるものとする)
func decrement(A int) int {
	last := A & 1
	return A>>1 | (last << 49)
}

// 集合Sの内容を昇順で出力する(スペース区切りで各要素の値を出力する)
func printSet(S int) {
	var cont []int
	for i := 0; i < 50; i++ {
		if S&(1<<i) != 0 {
			cont = append(cont, i)
		}
	}
	for i, val := range cont {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	M := getI()
	b := getInts(M)

	A, B := 0, 0
	for i := 0; i < N; i++ {
		A |= 1 << a[i]
	}
	for i := 0; i < M; i++ {
		B |= 1 << b[i]
	}

	com := getS()
	if com == "intersection" {
		printSet(intersection(A, B))
	} else if com == "union_set" {
		printSet(unionSet(A, B))
	} else if com == "symmetric_diff" {
		printSet(symmetricDiff(A, B))
	} else if com == "subtract" {
		x := getI()
		printSet(subtract(A, x))
	} else if com == "increment" {
		printSet(increment(A))
	} else if com == "decrement" {
		printSet(decrement(A))
	}
}
