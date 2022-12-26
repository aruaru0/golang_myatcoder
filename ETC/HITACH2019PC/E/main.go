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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	c := [3]int{0, 0, 0}
	ans := 1
	for _, v := range a {
		cnt := 0
		for i := 0; i < 3; i++ {
			if v == c[i] {
				cnt++
			}
		}
		ans *= cnt
		ans %= mod
		for i := 0; i < 3; i++ {
			if c[i] == v {
				c[i]++
				break
			}
		}
	}

	out(ans)
}
