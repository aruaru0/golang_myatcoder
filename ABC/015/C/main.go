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

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()

	T := make([][]int, N)
	for i := 0; i < N; i++ {
		T[i] = make([]int, K)
		for j := 0; j < K; j++ {
			T[i][j] = getInt()
		}
	}

	//out(T)

	n := 1
	for i := 0; i < N; i++ {
		n *= K
	}

	for i := 0; i < n; i++ {
		b := i
		a := make([]int, 0)
		for j := 0; j < N; j++ {
			a = append(a, b%K)
			b /= K
		}
		//out(a)
		ans := 0
		for j := 0; j < N; j++ {
			ans = ans ^ T[j][a[j]]
		}
		if ans == 0 {
			out("Found")
			return
		}
	}
	out("Nothing")
}
