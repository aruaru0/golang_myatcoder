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

func checkNext(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] <= x
	})
	return idx
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	s := getString()
	t := getString()
	p := make([][]int, 26)
	n := len(s)
	for i := 0; i < n; i++ {
		a := int(s[i] - 'a')
		p[a] = append(p[a], i)
	}

	ans := 0
	pos := 0
	for i := 0; i < len(t); i++ {
		x := int(t[i] - 'a')
		if len(p[x]) == 0 {
			out(-1)
			return
		}
		po := lowerBound(p[x], pos)
		// out(po, p[x], t[i:i+1])
		if po >= len(p[x]) {
			po = 0
			ans += n
		}
		pos = p[x][po] + 1
		// out(pos)
	}
	out(ans + pos)
}
