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
	N, K, S := getInt(), getInt(), getString()

	buy := 0
	rest := 0
	if N <= K {
		for _, e := range S {
			if rest > 0 {
				rest--
			} else {
				buy++
			}
			if e == '1' {
				rest++
			} else if e == '2' {
				rest += 2
			}
		}
		// out(rest, buy)
		// out(N, K)

		if rest >= buy {
			out(buy)
			return
		}
	}

	if K == N {
		out(buy)
		return
	}

	k := K % N
	buy += (K/N - 1) * (buy - rest)
	tmp := 0
	tmpRest := 0
	for i := 0; i < k; i++ {
		e := S[i]
		if tmpRest > 0 {
			tmpRest--
		} else {
			tmp++
		}
		if e == '1' {
			tmpRest++
		} else if e == '2' {
			tmpRest += 2
		}
	}
	// out(tmp, buy, rest)
	tmp = max(0, tmp-rest)
	out(buy + tmp)
}
