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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

var x, y []int

func veccross(p1, p2, p3 [2]int) int {
	p3[0] -= p1[0]
	p2[0] -= p1[0]
	p3[1] -= p1[1]
	p2[1] -= p1[1]
	return p3[0]*p2[1] - p2[0]*p3[1]
}

func ok(p []int) bool {
	PP := make([][2]int, 5)
	for i := 0; i < 5; i++ {
		PP[i] = [2]int{x[p[i]], y[p[i]]}
	}
	for i := 0; i < 5; i++ {
		c1 := veccross(PP[i], PP[(i+1)%5], PP[(i+2)%5])
		c2 := veccross(PP[i], PP[(i+1)%5], PP[(i+3)%5])
		c3 := veccross(PP[i], PP[(i+1)%5], PP[(i+4)%5])
		if c1*c2 >= 0 || c3*c2 >= 0 {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	x = make([]int, 5)
	y = make([]int, 5)
	for i := 0; i < 5; i++ {
		x[i], y[i] = getInt(), getInt()
	}

	a := []int{0, 1, 2, 3, 4}
	for {
		if ok(a) {
			out("YES")
			return
		}
		if !NextPermutation(sort.IntSlice(a)) {
			break
		}
	}
	out("NO")
}
