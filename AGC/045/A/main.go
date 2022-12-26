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

func bitCount(bits int) int {

	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

func rem(bs []int, num int) int {
	for i := 0; i < len(bs); i++ {
		if (bs[i] ^ num) < num {
			num ^= bs[i]
		}
	}
	return num
}

func solve() {
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	s := getString()

	b := make([]int, 0)
	for i := N - 1; i >= 0; i-- {
		if s[i] == '1' {
			if rem(b, a[i]) == 0 {
				continue
			}
			out(1)
			return
		} else {
			r := rem(b, a[i])
			if r == 0 {
				continue
			}
			b = append(b, r)
			for i := len(b) - 1; i > 0; i-- {
				if b[i] > b[i-1] {
					b[i], b[i-1] = b[i-1], b[i]
				}
			}
		}
	}
	out(0)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	T := getInt()
	for i := 0; i < T; i++ {
		solve()
	}
	//	out(1 << 62)
	// out(int(1e18))
}
