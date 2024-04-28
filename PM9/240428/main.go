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
	n := getI()
	a := getInts(n)
	b := getInts(n)

	root := make([]int, 0)
	ab := make([]int, 0)
	ai, bi := 0, 0
	for i := 0; i < n*2; i++ {
		var flag bool
		if ai <= n-1 && bi <= n-1 {
			flag = a[ai] < b[bi] // aの方が小さい = ガチャが先にある
		} else if ai > n-1 { // aが足りない
			flag = false
		} else {
			flag = true
		}
		if flag { // ガチャを回す場合
			root = append(root, a[ai])
			ab = append(ab, -1)
			ai++
		} else {
			root = append(root, b[bi])
			ab = append(ab, +1)
			bi++
		}
	}

	const inf = int(1e18)
	ans := inf
	coin := 0
	leftminus := 0
	for i := 0; i < n*2; i++ {
		coin += ab[i]
		ans = min(root[len(root)-1]+leftminus*2+root[len(root)-1]-root[i], ans)
		if i != n*2-1 && coin < 0 { // コインが不足する場合戻る必要がある
			leftminus += root[i+1] - root[i]
		}
	}

	out(ans)
}
