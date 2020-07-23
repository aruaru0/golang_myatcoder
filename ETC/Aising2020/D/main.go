package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

func popCount(X []byte) int {
	cnt := 0
	for _, v := range X {
		if v == '1' {
			cnt++
		}
	}
	return cnt
}

func g(x int) int {
	cnt := 0
	for x > 0 {
		cnt++
		x = x % bits.OnesCount(uint(x))
	}
	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	X := []byte(getString())

	popcount := popCount(X)
	popcount1 := popcount - 1
	popcount0 := popcount + 1

	mod1 := 0
	mod0 := 0
	for i := 0; i < N; i++ {
		if popcount1 != 0 {
			mod1 = mod1*2 + int(X[i]-'0')
			mod1 %= popcount1
		}
		mod0 = mod0*2 + int(X[i]-'0')
		mod0 %= popcount0
	}

	onePow2 := make([]int, 220000)
	zeroPow2 := make([]int, 220000)
	onePow2[0] = 1
	zeroPow2[0] = 1
	for i := 1; i < 220000; i++ {
		onePow2[i] = 1
		zeroPow2[i] = 1
		if popcount1 != 0 {
			onePow2[i] = onePow2[i-1] * 2 % popcount1
		}
		zeroPow2[i] = zeroPow2[i-1] * 2 % popcount0
	}

	for i := N - 1; i >= 0; i-- {
		if X[N-1-i] == '1' {
			if popcount1 != 0 {
				next := mod1
				next -= onePow2[i]
				next %= popcount1
				if next < 0 {
					next += popcount1
				}
				out(g(next) + 1)
				// out(f[next] + 1)
			} else {
				out(0)
			}
		} else {
			next := mod0
			next += zeroPow2[i]
			next %= popcount0
			out(g(next) + 1)
			// out(f[next] + 1)
		}
	}
}
