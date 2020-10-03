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

func f(n int, s string) {
	// check ooo
	o := 0
	o3 := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'o' {
			o++
		} else {
			if o >= 3 {
				o3 = true
			}
			o = 0
		}
	}
	if o == 3 {
		o3 = true
	}
	o2 := false
	for i := 2; i < len(s); i++ {
		x1, x2, x3 := s[i-2], s[i-1], s[i]
		if x1 == 'o' && x2 == 'o' && x3 == '-' {
			o2 = true
			break
		}
		if x1 == '-' && x2 == 'o' && x3 == 'o' {
			o2 = true
			break
		}
	}
	o1 := false
	for i := 3; i < len(s); i++ {
		x1, x2, x3, x4 := s[i-3], s[i-2], s[i-1], s[i]
		if x1 == '-' && x2 == 'o' && x3 == '-' && x4 == '-' {
			o1 = true
			break
		}
		if x1 == '-' && x2 == '-' && x3 == 'o' && x4 == '-' {
			o1 = true
			break
		}
	}

	o0 := false
	oflg := false
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'o' {
			if cnt%2 == 1 {
				o0 = true
				break
			}
			oflg = true
			continue
		}
		if oflg == true && s[i] == '-' {
			cnt++
		} else {
			cnt = 0
		}
		if s[i] == 'x' {
			cnt = 0
			oflg = false
		}
	}

	// out(o3, o2, o1, o0)
	// fmt.Print(num, ":")
	num++
	if o3 || o2 || o1 || o0 {
		out("O")
		return
	}
	out("X")
}

var num = 1

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	T := getInt()
	for i := 0; i < T; i++ {
		n, s := getInt(), getString()
		f(n, s)
	}
}
