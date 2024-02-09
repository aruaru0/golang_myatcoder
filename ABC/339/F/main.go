package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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

const mod = int(1e9 + 7)

func hash(x string) int {
	ret := 0
	for _, v := range x {
		e := int(v)
		ret = ret*37 + e
		ret %= mod
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	mod := big.NewInt(998244353)
	n := getI()
	c := map[string]int{}
	for i := 0; i < n; i++ {
		c[getS()]++
	}
	s := []string{}
	for v := range c {
		s = append(s, v)
	}
	sort.Strings(s)

	bn := []*big.Int{}
	bnc := []int{}
	for _, v := range s {
		b, _ := (&big.Int{}).SetString(v, 10)
		bn = append(bn, b)
		bnc = append(bnc, c[v])
	}

	m := make([]int, len(c))
	rm := map[int][]int{}
	for i, v := range bn {
		vv := int((&big.Int{}).Mod(v, mod).Int64())
		m[i] = vv
		rm[vv] = append(rm[vv], i)
	}

	ans := 0
	for i, a := range bn {
		for j := i; j < len(bn); j++ {
			b := bn[j]

			ab := (&big.Int{}).Mul(a, b)
			mab := int((&big.Int{}).Mod(ab, mod).Int64())
			idx := rm[mab]

			ok, ng := -1, len(idx)
			for abs(ok-ng) > 1 {
				mid := (ok + ng) / 2
				if bn[idx[mid]].Cmp(ab) <= 0 {
					ok = mid
				} else {
					ng = mid
				}
			}
			if ok >= 0 && bn[idx[ok]].Cmp(ab) == 0 {
				ans += bnc[i] * bnc[j] * bnc[idx[ok]]

				if i != j {
					ans += bnc[i] * bnc[j] * bnc[idx[ok]]
				}
			}
		}
	}

	out(ans)
}
