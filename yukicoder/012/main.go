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

func prime(n int) ([]int, []int) {
	a := make([]int, n+1)
	for i := 2; i*i <= n; i++ {
		if a[i] == 0 {
			for j := i * 2; j <= n; j += i {
				a[j] = 1
			}
		}
	}
	ret := make([]int, 0)
	for i := 2; i <= n; i++ {
		if a[i] == 0 {
			ret = append(ret, i)
		}
	}
	bits := make([]int, 0)
	for _, x := range ret {
		b := 0
		for x > 0 {
			b |= 1 << uint(x%10)
			x /= 10
		}
		bits = append(bits, b)
	}

	return ret, bits
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	mask := 1023
	check := 0
	minA := 9
	for i := 0; i < N; i++ {
		x := getInt()
		if x != 0 {
			minA = min(minA, x)
		}
		mask ^= 1 << uint(x)
		check |= 1 << uint(x)
	}

	p, b := prime(5000000)
	m := make([]int, len(p))
	for i := 0; i < len(b); i++ {
		m[i] = b[i] & mask
	}

	pos := -1
	p = append([]int{0}, p...)
	ans := -1
	for i := 0; i < len(m); i++ {
		if m[i] == 0 && p[i+1] >= minA {
			if pos == -1 {
				pos = i
			}
			continue
		} else {
			if pos != -1 {
				bits := 0
				for j := pos; j < i; j++ {
					bits |= b[j]
				}
				if bits == check {
					ans = max(ans, p[i+1]-p[pos]-2)
				}
			}
			pos = -1
		}
	}
	if pos != -1 {
		ans = max(ans, 5000000-p[pos]-1)
	}
	out(ans)
}
