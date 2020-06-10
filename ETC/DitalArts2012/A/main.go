package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func match(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i, v := range s {
		if t[i] == '*' {
			continue
		}
		if t[i] != byte(v) {
			return false
		}
	}
	return true
}

func main() {
	sc.Buffer([]byte{}, 1000000)
	sc.Split(bufio.ScanLines)
	S := getString()
	N := getInt()
	t := make([]string, N)
	for i := 0; i < N; i++ {
		t[i] = getString()
	}
	s := strings.Split(S, " ")
	for _, v := range s {
		ok := true
		for _, e := range t {
			if match(v, e) {
				ok = false
				break
			}
		}
		if ok == false {
			var o string
			for i := 0; i < len(v); i++ {
				o += "*"
			}
			fmt.Print(o, " ")
		} else {
			fmt.Print(v, " ")
		}
	}
	out()
}
