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

func f(a []int, s int) int {
	N := len(a)
	l := list.New()
	if s == 0 {
		l.PushBack(a[0])
		a = a[1:]
	} else {
		l.PushBack(a[N-1])
		a = a[:N-1]
	}
	ans := 0
	lpos, rpos := 0, len(a)-1
	for i := 0; i < len(a); i++ {
		f := l.Front().Value.(int)
		b := l.Back().Value.(int)
		diff := []int{
			abs(f - a[lpos]),
			abs(f - a[rpos]),
			abs(b - a[lpos]),
			abs(b - a[rpos])}
		sel := 0
		for j := 0; j < 4; j++ {
			if diff[sel] < diff[j] {
				sel = j
			}
		}
		switch sel {
		case 0:
			l.PushFront(a[lpos])
			lpos++
		case 1:
			l.PushFront(a[rpos])
			rpos--
		case 2:
			l.PushBack(a[lpos])
			lpos++
		case 3:
			l.PushBack(a[rpos])
			rpos--
		}
		ans += diff[sel]
	}

	return ans
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	sort.Ints(a)

	ans := max(f(a, 0), f(a, 1))
	out(ans)
}
