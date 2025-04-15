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
	l, r int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, k := getI(), getI()
	s := []byte(getS())
	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			if i != 0 {
				s[i-1] = '.'
			}
			if i+1 < n {
				s[i+1] = '.'
			}
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			cnt++
		}
	}

	x := k - cnt

	ps := make([]pair, 0)
	{
		for i := 0; i < n; {
			if s[i] == '?' {
				l := i
				for i < n && s[i] == '?' {
					i++
				}
				r := i
				ps = append(ps, pair{l, r})
			} else {
				i++
			}
		}
	}

	mx := 0
	for _, e := range ps {
		mx += (e.r - e.l + 1) / 2
	}

	if x == 0 {
		for _, e := range ps {
			for i := e.l; i < e.r; i++ {
				s[i] = '.'
			}
		}
	} else if x == mx {
		for _, e := range ps {
			if (e.r-e.l)%2 == 0 {
				continue
			}
			for i := 0; i < e.r-e.l; i++ {
				s[e.l+i] = "o."[i%2]
			}
		}
	}

	out(string(s))
}
