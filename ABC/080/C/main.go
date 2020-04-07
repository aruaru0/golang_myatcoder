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

func check(N int, a []int, f [][10]int, p [][11]int) int {
	ans := 0
	// out("-----")
	// out(a)
	for i := 0; i < N; i++ {
		// out(f[i])
		cnt := 0
		for j := 0; j < 10; j++ {
			if f[i][j] == 1 && a[j] == 1 {
				cnt++
			}
		}
		// out(cnt)
		ans += p[i][cnt]
	}
	// out(ans)
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	f := make([][10]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < 10; j++ {
			f[i][j] = getInt()
		}
	}
	p := make([][11]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j <= 10; j++ {
			p[i][j] = getInt()
		}
	}
	// out(N)
	// out(f)
	// out(p)
	b := 1 << 10
	ans := -1110000000
	for i := 1; i < b; i++ {
		a := make([]int, 10)
		n := i
		for j := 0; n > 0; j++ {
			a[j] = n % 2
			n /= 2
		}
		ans = max(ans, check(N, a, f, p))
	}
	out(ans)
}
