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
	N, M := getInt(), getInt()
	n := make([]int, N+1)
	for i := 0; i < M; i++ {
		a, b := getInt()-1, getInt()-1
		if a > b {
			a, b = b, a
		}
		n[a]++
		n[b]--
		// out(n)
	}
	// out(n)
	for i := 1; i <= N; i++ {
		n[i] += n[i-1]
	}
	// out(n)
	ans := "YES"
	for i := 0; i < N; i++ {
		if n[i]%2 == 1 {
			ans = "NO"
		}
	}
	out(ans)
}
