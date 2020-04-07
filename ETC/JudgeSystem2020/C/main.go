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

func check(a []int, a1, a2, a3 int) bool {
	x := make([][]int, 3)
	x[0] = a[:a1]
	x[1] = a[a1 : a1+a2]
	x[2] = a[a1+a2 : a1+a2+a3]
	ret := true
	m := min(a1, a2)
	// a1　a2
	for j := 0; j < m; j++ {
		if x[1][j] <= x[0][j] {
			ret = false
		}
	}
	// a2　a3
	m = min(a2, a3)
	for j := 0; j < m; j++ {
		if x[2][j] <= x[1][j] {
			ret = false
		}
	}
	// a1
	for j := 1; j < a1; j++ {
		if x[0][j] <= x[0][j-1] {
			ret = false
		}
	}
	// a2
	for j := 1; j < a2; j++ {
		if x[1][j] <= x[1][j-1] {
			ret = false
		}
	}
	// a3
	for j := 1; j < a3; j++ {
		if x[2][j] <= x[2][j-1] {
			ret = false
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)

	a1, a2, a3 := getInt(), getInt(), getInt()
	N := a1 + a2 + a3
	x := make([]int, N)
	for i := 1; i < N; i++ {
		x[i] = i
	}

	ret := 0
	f := check(x, a1, a2, a3)
	if f {
		ret++
	}
	cnt := 0
	for i := 1; NextPermutation(sort.IntSlice(x)); i++ {
		f = check(x, a1, a2, a3)
		if f {
			ret++
		}
		//	fmt.Println(i, x)
		cnt++
	}
	out(ret)
}
