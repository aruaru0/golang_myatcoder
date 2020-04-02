package main

import (
	"bufio"
	"fmt"
	"os"
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

// Pfs :　素因数分解し、スライスを作成
func Pfs(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// PfsMap : 素因数分解し、マップを作成
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

// 約数を列挙
func divisor(n int) []int {
	ret := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			// if i*i != n {
			ret = append(ret, n/i)
			// }
		}
	}
	return ret
}

func solve2(N int) int {
	x := divisor(N)
	n := len(x)
	a := x[n-1]
	b := x[n-2]
	return a + b - 2
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	ans2 := solve2(N)
	out(ans2)
}
