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

func calcPrime(N int) []int {
	prime := make([]int, N)
	for i := 0; i < N; i++ {
		prime[i] = i
	}
	for i := 2; i*i < N; i++ {
		if prime[i] == -1 {
			continue
		}
		for j := i * 2; j < N; j += i {
			prime[j] = -1
		}
	}
	return prime
}

func main() {
	sc.Split(bufio.ScanWords)
	X := getInt()

	p := calcPrime(1000000)

	for i := X; ; i++ {
		if p[i] != -1 {
			out(p[i])
			break
		}
	}
}
