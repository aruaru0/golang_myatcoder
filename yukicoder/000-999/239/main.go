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
	N := getInt()
	s := make([][]string, N)
	for i := 0; i < N; i++ {
		s[i] = make([]string, N)
		for j := 0; j < N; j++ {
			s[i][j] = getString()
		}
	}

	cnt := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s[i][j] == "nyanpass" {
				cnt[j]++
			}
		}
	}

	pos := -1
	for i := 0; i < N; i++ {
		if cnt[i] == N-1 {
			if pos == -1 {
				pos = i
			} else {
				out(-1)
				return
			}
		}
	}

	if pos != -1 {
		out(pos + 1)
		return
	}
	out(-1)
}
