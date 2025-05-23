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
	sc.Buffer([]byte{}, 1000000)
	S := getString()
	n := 0
	w := 0
	e := 0
	s := 0
	for _, v := range S {
		if v == 'N' {
			n++
		}
		if v == 'S' {
			s--
		}
		if v == 'W' {
			w++
		}
		if v == 'E' {
			e--
		}
	}
	if n == 0 && s == 0 && e != 0 && w != 0 {
		out("Yes")
		return
	}
	if n != 0 && s != 0 && e == 0 && w == 0 {
		out("Yes")
		return
	}
	if n != 0 && s != 0 && e != 0 && w != 0 {
		out("Yes")
		return
	}
	out("No")
}
