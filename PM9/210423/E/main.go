package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

func f(n int) int {
	cnt := 0
	for n > 0 {
		k := bits.OnesCount(uint(n))
		n %= k
		cnt++
	}
	return cnt
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x := getS()

	tot := 0
	for i := 0; i < N; i++ {
		if x[i] == '1' {
			tot++
		}
	}

	// １の数はtot-1, tot-1の２パターンしかない
	cnt0, cnt1 := tot-1, tot+1
	if cnt0 == 0 {
		cnt0 = 1
	}

	// あらかじめ余りを計算しておく（２パターン）
	p0 := 0
	p1 := 0
	for i := 0; i < N; i++ {
		v := int(x[i] - '0')
		p0 = (p0*2 + v) % cnt0
		p1 = (p1*2 + v) % cnt1
	}

	t0, t1 := 1, 1
	o := make([]int, 0)
	for i := N; i > 0; i-- {
		v := int(x[i-1] - '0')

		if v == 0 {
			// 0 -> 1になった時は、該当桁の余りを足して計算
			ans := (p1 + t1) % cnt1
			o = append(o, f(ans)+1)
		} else {
			if tot-1 == 0 { // mod 0のときは０を出力
				o = append(o, 0)
				continue
			}
			// 1 -> 0になった時は、該当桁を引いて計算
			ans := (p0 - t0) % cnt0
			if ans < 0 { // 負になった時の処理
				ans += cnt0
			}
			o = append(o, f(ans)+1)
		}

		t0 = t0 * 2 % cnt0
		t1 = t1 * 2 % cnt1
	}

	for i := len(o) - 1; i >= 0; i-- {
		out(o[i])
	}
}
