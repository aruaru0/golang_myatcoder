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

type LR struct {
	l, r, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	lr := make([]LR, N)
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
		l, r := 0, 0
		for _, e := range s[i] {
			if e == '(' {
				l++
			} else {
				if l > 0 {
					l--
				} else {
					r++
				}
			}
		}
		lr[i] = LR{l, r, i}
	}

	sort.Slice(lr, func(i, j int) bool {
		di := lr[i].l - lr[i].r
		dj := lr[j].l - lr[j].r
		if di > 0 && dj > 0 { // 両方とも　（　が多い場合には、　）の少ない順
			return lr[i].r < lr[j].r
		} else if di <= 0 && dj <= 0 { // 両方とも　）　が多い場合には　（の多い順
			return lr[i].l > lr[j].l
		}
		return di > dj // 上記以外は、（の多いほうを手前に配置
	})

	l := 0
	ok := true
	for i := 0; i < N; i++ {
		x := s[lr[i].idx]
		for _, e := range x {
			if e == '(' {
				l++
			} else {
				if l > 0 {
					l--
				} else {
					ok = false
				}
			}
		}
	}

	if l == 0 && ok {
		out("Yes")
	} else {
		out("No")
	}
}
