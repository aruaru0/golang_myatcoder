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

func check(s, t string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			continue
		}
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	s := getString()
	t := getString()

	x := make([]string, 0)
	for i := len(s) - len(t); i >= 0; i-- {
		flg := check(s[i:i+len(t)], t)
		if flg {
			o := []byte(s[0:i] + t + s[i+len(t):])
			for j := 0; j < len(o); j++ {
				if o[j] == '?' {
					o[j] = 'a'
				}
			}
			x = append(x, string(o))
		}
	}

	sort.Strings(x)
	if len(x) > 0 {
		out(x[0])
	} else {
		out("UNRESTORABLE")
	}
}
