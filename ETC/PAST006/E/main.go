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
	N := getI()
	S := getS()
	l := list.New()

	for i := 0; i < N; i++ {
		switch S[i] {
		case 'L':
			l.PushFront(i + 1)
		case 'R':
			l.PushBack(i + 1)
		case 'A':
			if l.Len() <= 0 {
				out("ERROR")
			} else {
				p := l.Front()
				out(p.Value)
				l.Remove(p)
			}
		case 'B':
			if l.Len() <= 1 {
				out("ERROR")
			} else {
				p := l.Front()
				p = p.Next()
				out(p.Value)
				l.Remove(p)
			}
		case 'C':
			if l.Len() <= 2 {
				out("ERROR")
			} else {
				p := l.Front()
				p = p.Next()
				p = p.Next()
				out(p.Value)
				l.Remove(p)
			}
		case 'D':
			if l.Len() <= 0 {
				out("ERROR")
			} else {
				p := l.Back()
				out(p.Value)
				l.Remove(p)
			}
		case 'E':
			if l.Len() <= 1 {
				out("ERROR")
			} else {
				p := l.Back()
				p = p.Prev()
				out(p.Value)
				l.Remove(p)
			}
		case 'F':
			if l.Len() <= 2 {
				out("ERROR")
			} else {
				p := l.Back()
				p = p.Prev()
				p = p.Prev()
				out(p.Value)
				l.Remove(p)
			}
		}
	}

}
