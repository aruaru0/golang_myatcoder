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

// NextPermutation generates the next permutation of the
// sortable collection x in lexical order.  It returns false
// if the permutations are exhausted.
//
// Knuth, Donald (2011), "Section 7.2.1.2: Generating All Permutations",
// The Art of Computer Programming, volume 4A.
// ※NextPermutationは辞書順で次を返す
func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

type dat struct {
	l, r, n int
}

var N int
var d []dat

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	s := getS()

	d = make([]dat, N)
	l, r := 0, 0
	n := 0
	tot := 0
	for i := 0; i < N; i++ {
		if s[i] == '(' {
			l++
		} else {
			if l > 0 {
				l--
				n++
			} else {
				r++
			}
		}
		d[i] = dat{l, r, n}
		tot += n
	}
	sort.Slice(d, func(i, j int) bool {
		mi0 := min(-d[i].r, -d[i].r+d[i].l-d[j].r)
		ma0 := max(-d[i].r, -d[i].r+d[i].l-d[j].r)
		mi1 := min(-d[j].r, -d[j].r+d[j].l-d[i].r)
		ma1 := max(-d[j].r, -d[j].r+d[j].l-d[i].r)
		if mi0 == mi1 {
			return ma0 > ma1
		}
		return mi0 > mi1
	})

	// out(d)
	cur := 0
	for _, e := range d {
		tot += min(cur, e.r)
		cur -= min(cur, e.r)
		cur += e.l
	}
	out(tot * 2)

}
