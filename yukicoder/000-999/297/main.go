package main

import (
	"bufio"
	"fmt"
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

func makeMax(a []int, op []int) int {
	l := len(a) - len(op)
	b := 0
	for i := 0; i < l; i++ {
		b = b*10 + a[i]
	}
	a = a[l:]
	// out(b, a, op)
	// max
	ma := b
	for i := 0; i < len(op); i++ {
		if op[i] == 1 {
			ma += a[i]
		} else {
			ma -= a[i]
		}
	}
	return ma
}

func makeMin(a, op []int) int {
	flg := false
	for _, e := range op {
		if e == -1 {
			flg = true
		}
	}

	mi := 0

	if flg {
		l := len(a) - len(op)
		b := 0
		for i := 0; i < l; i++ {
			b = b*10 + a[i]
		}
		a = a[l:]
		// min
		mi = a[len(op)-1]
		// out(b, a, op)
		pos := len(a) - 2
		for i := 0; i < len(op)-1; i++ {
			if op[i] == 1 {
				mi += a[pos]
			} else {
				mi -= a[pos]
			}
			pos--
		}
		if op[len(op)-1] == 1 {
			mi += b
		} else {
			mi -= b
		}
	} else {
		sort.Ints(a)
		l := len(op) + 1
		b := make([]int, l)
		for i := 0; i < len(a); i++ {
			b[i%l] = b[i%l]*10 + a[i]
		}
		mi = 0
		for _, e := range b {
			mi += e
		}
	}

	return mi
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	op := make([]int, 0)
	a := make([]int, 0)
	for i := 0; i < N; i++ {
		c := getS()
		if c == "+" {
			op = append(op, 1)
			continue
		}
		if c == "-" {
			op = append(op, -1)
			continue
		}
		a = append(a, int(c[0]-'0'))
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(op, func(i, j int) bool {
		return op[i] > op[j]
	})

	out(makeMax(a, op), makeMin(a, op))
}
