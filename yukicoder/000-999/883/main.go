package main

import (
	"bufio"
	"fmt"
	"math"
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
	N, K := getInt(), getInt()

	m := int(math.Ceil(math.Sqrt(float64(N))))
	if N >= K*K {
		m = (N + K - 1) / K
	}

	a := make([][]byte, m)
	b := make([]byte, m)
	for i := 0; i < m; i++ {
		if i < K {
			b[i] = '#'
		} else {
			b[i] = '.'
		}
	}

	for i := 0; i < m; i++ {
		a[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			if N != 0 {
				a[i][(i+j)%m] = b[j]
				if b[j] == '#' {
					N--
				}
			} else {
				a[i][(i+j)%m] = '.'
			}
		}
	}
	out(m)
	for i := 0; i < m; i++ {
		out(string(a[i]))
	}
}
