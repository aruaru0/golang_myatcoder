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

func reverse(S string) string {
	s := []byte(S)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return string(s)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := getString()

	s = reverse(s)
	d0 := reverse("dream")
	d1 := reverse("dreamer")
	d2 := reverse("erase")
	d3 := reverse("eraser")

	pos := 0
	// out(s)
	for pos != len(s) {
		if strings.HasPrefix(s[pos:], d0) {
			// out(d0)
			pos += len(d0)
		} else if strings.HasPrefix(s[pos:], d1) {
			// out(d1)
			pos += len(d1)
		} else if strings.HasPrefix(s[pos:], d2) {
			// out(d2)
			pos += len(d2)
		} else if strings.HasPrefix(s[pos:], d3) {
			// out(d3)
			pos += len(d3)
		} else {
			out("NO")
			return
		}
	}
	out("YES")
}
