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

// ｎ進数を計算（-2進数などに対応）
func calcN(x, n int) []int {
	ret := make([]int, 0)
	for x != 0 {
		r := x % n
		if r < 0 {
			r += (-n)
		}
		x = (x - r) / n
		ret = append(ret, r)
	}
	if len(ret) == 0 {
		ret = append(ret, 0)
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	a := make([]byte, 0)
	for N != 0 {
		r := N % 2
		if r < 0 {
			r += 2
		}
		N = (N - r) / (-2)

		a = append(a, byte(r))
	}
	// out(a)
	if len(a) == 0 {
		out(0)
		return
	}
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Print(string(a[i] + '0'))
	}
	out()
}
