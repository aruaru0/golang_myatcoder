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

	N, C := getInt(), getInt()
	T := make([][]int, C)
	for i := 0; i < C; i++ {
		T[i] = make([]int, 100100)
	}
	for i := 0; i < N; i++ {
		s, t, c := getInt(), getInt(), getInt()-1
		T[c][s]++
		T[c][t+1]--
	}

	for c := 0; c < C; c++ {
		for i := 1; i < 100100; i++ {
			T[c][i] += T[c][i-1]
		}
	}

	ans := 0
	for i := 1; i < 100100; i++ {
		cnt := 0
		for c := 0; c < C; c++ {
			if T[c][i] != 0 {
				cnt++
			}
		}
		ans = max(ans, cnt)
	}

	out(ans)
}
