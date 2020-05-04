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
	D, N := getInt(), getInt()

	x := 1
	for i := 0; i < D; i++ {
		x *= 100
	}

	y := 0
	cnt := 0
	for {
		y += x
		if y%(x*100) == 0 {
			continue
		}
		cnt++
		if cnt == N {
			break
		}
	}

	// for i := 0; i < N; i++ {
	// 	y += x
	// }

	// if x*100 == y {
	// 	y += x
	// }
	out(y)
}
