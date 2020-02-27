package main

import (
	"bufio"
	"fmt"
	"math"
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

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N)
	s := make([]int, N)
	sum := 0
	for i := 0; i < N; i++ {
		a[i] = getInt()
		sum += a[i]
		s[i] = sum
	}

	ans := math.MaxInt64
	for i := 0; i < N-1; i++ {
		A := s[i]
		B := s[N-1] - s[i]
		diff := asub(A, B)
		if diff < ans {
			ans = diff
		}
		//		out(A, B, diff)
	}
	out(ans)

}
