package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
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

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	p := make([]int, N)
	q := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getInt()
	}
	for i := 0; i < N; i++ {
		q[i] = getInt()
	}

	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i + 1
	}

	flg := true
	ppos := -1
	qpos := -1
	idx := 0
	for flg {
		// out(a)
		ok := true
		for j := 0; j < N; j++ {
			if a[j] != p[j] {
				ok = false
			}
		}
		if ok {
			ppos = idx
		}
		ok = true
		for j := 0; j < N; j++ {
			if a[j] != q[j] {
				ok = false
			}
		}
		if ok {
			qpos = idx
		}
		idx++
		flg = NextPermutation(sort.IntSlice(a))
	}

	out(abs(ppos - qpos))
}
