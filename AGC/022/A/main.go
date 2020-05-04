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

// NextPermutationを使わない場合
func f(s string) {
	pos := -1
	m := make([]int, 26)
	v := byte(0)
L0:
	for i := len(s) - 1; i >= 0; i-- {
		c := int(s[i] - 'a')
		// out(c, m)
		for j := c; j < 26; j++ {
			if m[j] == 1 {
				v = byte(j + 'a')
				pos = i
				break L0
			}
		}
		m[c]++
	}

	if pos == -1 {
		out(-1)
		return
	}

	out(s[:pos] + string(v))

	return
}

func f2(s string) {
	if s == "zyxwvutsrqponmlkjihgfedcba" {
		out(-1)
		return
	}
	x := make([]int, 26)
	for i := 0; i < 26; i++ {
		x[i] = int(s[i])
	}

	NextPermutation(sort.IntSlice(x))

	for i := 0; i < 26; i++ {
		c := byte(x[i])
		if s[i] == c {
			fmt.Printf("%c", c)
		} else {
			fmt.Printf("%c", c)
			break
		}
	}
	out("")
}

func main() {
	sc.Split(bufio.ScanWords)
	s := getString()
	m := make([]int, 26)
	for _, v := range s {
		m[int(v-'a')]++
	}
	c := -1
	for i := 0; i < 26; i++ {
		if m[i] == 0 {
			c = i
			break
		}
	}
	if c == -1 {
		f2(s)
		return
	}
	out(s + string(byte(c+'a')))
}
