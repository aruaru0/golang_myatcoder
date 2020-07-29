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
	A, B, C := getInt(), getInt(), getInt()
	K := getInt()
	k := 1
	for i := 0; i < K; i++ {
		k *= 3
	}

	if A < B && B < C {
		out("Yes")
		return
	}

	for i := 0; i < k; i++ {
		n := i
		a, b, c := A, B, C
		for j := 0; j < K; j++ {
			sel := n % 3
			switch sel {
			case 0:
				a *= 2
			case 1:
				b *= 2
			case 2:
				c *= 2
			}
			n /= 3
			if a < b && b < c {
				out("Yes")
				return
			}
		}
	}
	out("No")
}
