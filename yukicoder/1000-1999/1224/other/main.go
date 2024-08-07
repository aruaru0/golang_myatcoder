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

func PfsMap(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2] = pfs[2] + 1
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i] = pfs[i] + 1
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = pfs[n] + 1
	}

	return pfs
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b := getInt(), getInt()

	pa := PfsMap(a)
	pb := PfsMap(b)

	for n, e := range pb {
		if n == 2 {
			pb[n] = 0
		}
		if n == 5 {
			pb[n] = 0
		}
		if pa[n] != 0 {
			pb[n] = max(0, e-pa[n])
		}
	}

	cnt := 0
	for _, e := range pb {
		cnt += e
	}
	if cnt != 0 {
		out("Yes")
		return
	}
	out("No")
}
