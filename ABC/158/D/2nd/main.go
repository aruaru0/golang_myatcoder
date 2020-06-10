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
	sc.Buffer([]byte{}, 1000000)

	S := getString()
	Q := getInt()
	s := make([]byte, len(S)+2*Q)
	for i := 0; i < len(S); i++ {
		s[i+Q] = S[i]
	}
	l := Q - 1
	r := len(S) + Q
	flg := false
	for i := 0; i < Q; i++ {
		t := getInt()
		if t == 1 {
			flg = !flg
		} else {
			f, c := getInt(), getString()[0]
			if f == 1 {
				if !flg {
					s[l] = byte(c)
					l--
				} else {
					s[r] = byte(c)
					r++
				}
			} else {
				if flg {
					s[l] = byte(c)
					l--
				} else {
					s[r] = byte(c)
					r++
				}

			}
		}
	}

	l++
	r--

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	if !flg {
		for i := l; i <= r; i++ {
			fmt.Fprint(w, string(s[i]))
		}
		fmt.Fprintln(w)
	} else {
		for i := r; i >= l; i-- {
			fmt.Fprint(w, string(s[i]))
		}
		fmt.Fprintln(w)
	}
}
