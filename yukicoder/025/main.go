package main

import (
	"bufio"
	"fmt"
	"math/big"
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

func pfs(n int) []int {
	ret := make([]int, 0)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			ret = append(ret, i)
			n /= i
		}
	}
	if n != 1 {
		ret = append(ret, n)
	}
	return ret
}
func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()

	if N%M == 0 {
		s := strconv.FormatInt(int64(N/M), 10)
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] != '0' {
				out(string(s[i]))
				return
			}
		}
	}

	X := M
	for {
		if X%2 == 0 {
			X /= 2
		} else if X%5 == 0 {
			X /= 5
		} else {
			break
		}
	}
	if X != 1 {
		out(-1)
		return
	}

	n := big.NewInt(int64(N))
	m := big.NewInt(int64(M))
	ten := big.NewInt(10)
	t := big.NewInt(0)

	for t.Mod(n, m).Int64() != 0 {
		n = n.Mod(n, m)
		n = n.Mul(n, ten)
	}
	x := n.Div(n, m).Int64()
	s := strconv.FormatInt(x, 10)
	out(s[:1])
}
