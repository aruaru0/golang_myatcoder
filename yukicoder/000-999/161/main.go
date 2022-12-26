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
	G, C, P := getInt(), getInt(), getInt()
	s := getString()
	g, c, p := 0, 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'G':
			g++
		case 'C':
			c++
		case 'P':
			p++
		}
	}

	ans := 0
	// win
	d := min(G, c)
	G -= d
	c -= d
	ans += d * 3
	d = min(C, p)
	C -= d
	p -= d
	ans += d * 3
	d = min(P, g)
	P -= d
	g -= d
	ans += d * 3
	// draw
	ans += min(G, g) + min(C, c) + min(P, p)

	out(ans)
}
