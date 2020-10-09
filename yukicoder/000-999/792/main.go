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

func main() {
	sc.Split(bufio.ScanWords)
	n := getInt()
	a := make([][]int, 0)
	s := 0
	for i := 0; i < 1<<n; i++ {
		x := getInts(n + 1)
		s += x[n]
		if x[n] == 1 {
			a = append(a, x)
		}
	}
	if s == 0 {
		out("A=⊥")
		return
	}
	if s == 1<<n {
		out("A=⊤")
		return
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprint(w, "A=")
	for i := 0; i < len(a); i++ {
		fmt.Fprint(w, "(")
		flg := a[i][n]
		for j := 0; j < n; j++ {
			if flg == 0 && a[i][j] == 0 {
				fmt.Fprintf(w, "P_%d", j+1)
			}
			if flg == 0 && a[i][j] == 1 {
				fmt.Fprintf(w, "¬P_%d", j+1)
			}
			if flg == 1 && a[i][j] == 0 {
				fmt.Fprintf(w, "¬P_%d", j+1)
			}
			if flg == 1 && a[i][j] == 1 {
				fmt.Fprintf(w, "P_%d", j+1)
			}
			if j != n-1 {
				fmt.Fprint(w, "∧")
			}
		}
		fmt.Fprint(w, ")")
		if i != len(a)-1 {
			fmt.Fprint(w, "∨")
		}
	}
	fmt.Fprintln(w)
}
