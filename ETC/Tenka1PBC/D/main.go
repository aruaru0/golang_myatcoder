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

func check(N int) int {
	for i := 0; i < 500; i++ {
		if i*(i-1)/2 == N {
			return i
		}
	}
	return -1
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	n := check(N)
	if n == -1 {
		out("No")
		return
	}
	m := make([][]int, n)
	cnt := 1
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// out(m)
			m[i][j] = cnt
			m[j][i] = cnt
			cnt++
		}
	}
	out("Yes")
	out(n)
	for i := 0; i < n; i++ {
		fmt.Print(n-1, " ")
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			fmt.Print(m[i][j], " ")
		}
		out()
	}
}
