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

func prime(x int) []int {
	d := make([]int, x+1)
	ret := make([]int, 0)
	for i := 2; i <= x; i++ {
		if d[i] == 0 {
			ret = append(ret, i)
			for j := i; j <= x; j += i {
				d[j] = 1
			}
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	p := prime(5000000)
	p = append([]int{0}, p...)
	p = append(p, 5000000)
	m := make([]int, len(p))
	for i, v := range p {
		for v > 0 {
			m[i] |= 1 << (v % 10)
			v /= 10
		}
	}
	mask := 0
	N := getInt()
	for i := 0; i < N; i++ {
		mask |= 1 << getInt()
	}
	if mask == 1023 {
		out(4999999)
		return
	}
	ret := -1
	for i := 1; i < len(p)-1; i++ {
		x := 0
		for j := 0; j < len(p)-i; j++ {
			x |= m[i+j]
			// fmt.Printf("%b %b\n", mask, x)
			if x&(1023^mask) != 0 {
				break
			}
			if x == mask {
				// out(i, j, p[i:i+j+1], m[j])
				ret = max(ret, p[i+j+1]-p[i-1]-2)
			}
		}
	}
	out(ret)
}
