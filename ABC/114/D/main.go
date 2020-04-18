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

func divisor(n int) []int {
	ret := []int{}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			// if i*i != n {
			ret = append(ret, n/i)
			// }
		}
	}
	return ret
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()

	a := make([]int, 0)
	for i := 1; i <= N; i++ {
		x := Pfs(i)
		// out(x, i)
		a = append(a, x...)
	}
	sort.Ints(a)
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	cnt74 := 0
	cnt24 := 0
	cnt14 := 0
	cnt4 := 0
	cnt2 := 0
	for _, v := range m {
		if v >= 74 {
			cnt74++
		}
		if v >= 24 {
			cnt24++
		}
		if v >= 14 {
			cnt14++
		}
		if v >= 4 {
			cnt4++
		}
		if v >= 2 {
			cnt2++
		}
	}

	// out(cnt74, ":", cnt24, cnt2, ":", cnt14, cnt4, ":", cnt4, cnt2, ":", len(m)-cnt2)

	ans := cnt74 + cnt24*(cnt2-1) + cnt14*(cnt4-1) + cnt4*(cnt4-1)*(cnt2-2)/2
	out(ans)
}
