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

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	N, Q := getInt(), getInt()
	a := make([]int, N)
	s := make([]int, Q)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	for i := 0; i < Q; i++ {
		s[i] = getInt()
	}

	// 一度計算したものをメモ
	m := make(map[int]int)
	for i := 0; i < Q; i++ {
		v := s[i]
		x := make(map[int]int)
		for j := 0; j < N; j++ {
			c, ok := m[v]
			if ok {
				for v := range x {
					m[v] = c
				}
				out(c)
				break
			}
			x[v] = 1
			v = GCD(v, a[j])
			// out("v", v, x)
			if v == 1 {
				for v := range x {
					m[v] = j + 1
				}
				out(j + 1)
				// out(m)
				break
			} else if j == N-1 {
				// out(m, x)
				for z := range x {
					m[z] = v
				}
				// out(m)
				out(v)
				// out(m)
				break
			}
		}
	}
}
