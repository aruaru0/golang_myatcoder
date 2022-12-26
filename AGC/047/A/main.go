package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

func convf(s string) int {
	x := strings.Split(s, ".")
	if len(x) == 1 {
		x = append(x, "0")
	}
	x[1] += "0000000000"
	x[1] = x[1][:9]
	a, _ := strconv.Atoi(x[0])
	b, _ := strconv.Atoi(x[1])
	return a*int(1e9) + b
}

const div = int(1e9)

func pfs25(n int) (int, int) {
	ret2, ret5 := 0, 0
	for n%2 == 0 {
		ret2++
		n /= 2
	}
	for n%5 == 0 {
		ret5++
		n /= 5
	}
	return ret2, ret5
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	a := make([]int, N)
	a2 := make([]int, 0)
	a5 := make([]int, 0)
	var X [40][40]int
	minus := 0
	for i := 0; i < N; i++ {
		s := getString()
		a[i] = convf(s)
		x, y := pfs25(a[i])
		if x*2 >= 18 && y*2 >= 18 {
			minus++
		}
		x = min(18, x)
		y = min(18, y)
		a2 = append(a2, x)
		a5 = append(a5, y)
		X[x][y]++
	}
	// out(a2)
	// out(a5)
	// out(minus)

	cnt := 0
	for i := 0; i < len(a2); i++ {
		x := 18 - a2[i]
		y := 18 - a5[i]
		// out(a[i], a2[i], a5[i], x, y)
		for j := x; j <= 18; j++ {
			for k := y; k <= 18; k++ {
				// if X[j][k] != 0 {
				// 	out(X[j][k], j, k)
				// }
				cnt += X[j][k]
			}
		}
	}

	out((cnt - minus) / 2)
}
