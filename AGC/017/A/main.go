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

var C [52][52]int // C[n][k] -> nCk

func combTable(N int) {
	C[0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if j == 0 || j == i {
				C[i][j] = 1
			} else {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N, P := getInt(), getInt()
	d1 := 0
	d2 := 0
	for i := 0; i < N; i++ {
		a := getInt()
		if a%2 == 0 {
			d2++
		} else {
			d1++
		}
	}

	off := 0
	if P == 1 {
		off = 1
	}

	combTable(51)

	ans := 0
	for i := 0; i <= d2; i++ {
		for j := off; j <= d1; j += 2 {
			// out(i, j, "C", d2, i, "C", d1, j)
			ans += C[d2][i] * C[d1][j]
		}
	}
	out(ans)
}
