package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

func solve(a, b, c, d int) int {
	return 0
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	a, b := getI(), getI()
	c, d := getI(), getI()

	if a == c && b == d {
		out(0)
		return
	}

	if a+b == c+d || a-b == c-d {
		out(1)
		return
	}
	if abs(a-c)+abs(b-d) <= 3 {
		out(1)
		return
	}

	if abs(a-c)+abs(b-d) <= 6 {
		out(2)
		return
	}

	if a%2 == b%2 && c%2 == d%2 {
		out(2)
		return
	}
	if a%2 != b%2 && c%2 != d%2 {
		out(2)
		return
	}

	for i := -3; i <= 3; i++ {
		for j := -3; j <= 3; j++ {
			if abs(i)+abs(j) <= 3 {
				cc := c + i
				dd := d + j
				if a+b == cc+dd || a-b == cc-dd {
					out(2)
					return
				}
			}
		}
	}

	out(3)
}
