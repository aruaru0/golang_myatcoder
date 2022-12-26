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

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func test() {
	L := getInt()
	l := L / 4
	cnt := 0
	for a := 1; a < l; a++ {
		for b := a + 1; b < l; b++ {
			for c := b + 1; c < l; c++ {
				x := GCD(a, GCD(b, c))
				if x == 1 && a+b+c <= l && a*a+b*b == c*c {
					out(a, b, c)
					cnt++
				}
			}
		}
	}
	out(cnt)
}

func main() {
	sc.Split(bufio.ScanWords)
	L := getInt()
	l := L / 4
	cnt := 0
	for n := 1; n*n <= l; n++ {
		for m := n; m*m <= l; m++ {
			if GCD(m, n) != 1 {
				continue
			}
			if (m-n)%2 != 1 {
				continue
			}
			a := m*m - n*n
			b := 2 * m * n
			c := m*m + n*n
			if a+b+c <= l {
				// out(a, b, c)
				cnt++
			}
		}
	}

	out(cnt)
}
