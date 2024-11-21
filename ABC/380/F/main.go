package main

import (
	"bufio"
	"container/list"
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
	n, m, l := getI(), getI(), getI()
	as := getInts(n)
	bs := getInts(m)
	cs := getInts(l)
	abcs := []int{0}
	abcs = append(abcs, as...)
	abcs = append(abcs, bs...)
	abcs = append(abcs, cs...)

	st := [13]int{}
	for i := 1; i <= n; i++ {
		st[i] = 1
	}
	for i := n + 1; i <= n+m; i++ {
		st[i] = 2
	}
	mp := make(map[[13]int][]int)
	mpi := make(map[[13]int]int)
	mpir := make([][13]int, 1594323+10)
	used := make(map[[13]int]bool)
	gi := func(t [13]int) int {
		if v, ok := mpi[t]; ok {
			return v
		}
		mpi[t] = len(mpi)
		mpir[len(mpi)-1] = t
		return len(mpi) - 1
	}
	gi(st)
	q := list.New()
	put := func(t [13]int) {
		if _, ok := used[t]; ok {
			return
		}
		used[t] = true
		q.PushBack(t)
	}
	q.PushBack(st)
	e := q.Front()
	for e != nil {
		t := e.Value.([13]int)
		hv := []int{}
		bd := []int{}
		for i := 1; i <= n+m+l; i++ {
			if t[i] == t[0]+1 {
				hv = append(hv, i)
			}
			if t[i] == 0 {
				bd = append(bd, i)
			}
		}
		for _, i := range hv {
			nst := t
			nst[0] ^= 1
			nst[i] = 0
			mp[t] = append(mp[t], gi(nst))
			put(nst)
			for _, j := range bd {
				if abcs[i] > abcs[j] {
					nst := t
					nst[0] ^= 1
					nst[i] = 0
					nst[j] = t[0] + 1
					mp[t] = append(mp[t], gi(nst))
					put(nst)
				}
			}
		}
		e = e.Next()
	}

	var dfs func(v int) bool

	used2 := make([]int, len(mpi)+10)

	dfs = func(v int) bool {
		if used2[v] != 0 {
			if used2[v] == 1 {
				return true
			} else {
				return false
			}
		}

		t := mpir[v]
		if t[0] == 0 {
			r := false
			for _, nv := range mp[t] {
				if dfs(nv) {
					r = true
				}
			}
			if r {
				used2[v] = 1
			} else {
				used2[v] = 2
			}
			return r
		} else {
			r := true
			for _, nv := range mp[t] {
				if !dfs(nv) {
					r = false
				}
			}
			if r {
				used2[v] = 1
			} else {
				used2[v] = 2
			}
			return r
		}
	}
	r := dfs(mpi[st])
	if r {
		out("Takahashi")
	} else {
		out("Aoki")
	}

}
