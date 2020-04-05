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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K, C := getInt(), getInt(), getInt()
	s := []byte(getString())

	a := make([][]int, 2)
	sum := [2]int{0, 0}
	for ri := 0; ri < 2; ri++ {
		a[ri] = make([]int, N)
		c := 0
		cnt := 1
		for i := 0; i < N; i++ {
			// out(i, c, s[i])
			if c <= 0 && s[i] == 'o' {
				a[ri][i] = cnt
				cnt++
				c = C + 1
			}
			c--
		}
		sum[ri] = cnt - 1
		for i := 0; i < N/2; i++ {
			s[i], s[N-1-i] = s[N-1-i], s[i]
		}
	}

	// out(a[0])

	for i := 0; i < N/2; i++ {
		j := N - 1 - i
		a[1][i], a[1][j] = a[1][j], a[1][i]
	}

	// out(a[1])
	// out(sum)
	if sum[0] > K && sum[1] > K {
		return
	}

	for i := 0; i < N; i++ {
		if a[0][i] == 0 || a[1][i] == 0 {
			continue
		} else {
			out(i + 1)
		}
	}

}
