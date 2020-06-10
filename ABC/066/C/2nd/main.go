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
	sc.Split(bufio.ScanWords)
	n := getInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = getInt()
	}

	b := make([]int, n*2)
	l := n - 1
	r := n
	rot := false
	for i := 0; i < n; i++ {
		if !rot {
			b[r] = a[i]
			r++
		} else {
			b[l] = a[i]
			l--
		}
		rot = !rot
	}

	// out(b)
	// out(l, r)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	l++
	r--
	for i := 0; i < n; i++ {
		if !rot {
			fmt.Fprint(w, b[l], " ")
			l++
		} else {
			fmt.Fprint(w, b[r], " ")
			r--
		}
	}
	fmt.Fprintln(w)
}
