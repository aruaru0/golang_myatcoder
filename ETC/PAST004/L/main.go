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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, Q := getI(), getI()
	h := getInts(N)

	m := make(map[int]int)
	var mmap = func() {
		for i := 0; i < N-1; i++ {
			if i%2 == 0 {
				m[h[i]-h[i+1]]++
			} else {
				m[h[i+1]-h[i]]++
			}
		}
	}
	mmap()
	add := 0
	sub := 0
	for i := 0; i < Q; i++ {
		t := getI()
		switch t {
		case 1:
			add += getI()
		case 2:
			sub += getI()
		case 3:
			u, v := getI()-1, getI()
			if u%2 == 0 {
				if u+1 != N {
					m[h[u]-h[u+1]]--
				}
				if u-1 >= 0 {
					m[h[u]-h[u-1]]--
				}
				h[u] += v
				if u+1 != N {
					m[h[u]-h[u+1]]++
				}
				if u-1 >= 0 {
					m[h[u]-h[u-1]]++
				}
			} else {
				if u+1 != N {
					m[h[u+1]-h[u]]--
				}
				if u-1 >= 0 {
					m[h[u-1]-h[u]]--
				}
				h[u] += v
				if u+1 != N {
					m[h[u+1]-h[u]]++
				}
				if u-1 >= 0 {
					m[h[u-1]-h[u]]++
				}
			}
		}
		out(m[sub-add])
	}
}
