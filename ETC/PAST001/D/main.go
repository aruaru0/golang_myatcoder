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
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	a := make([]int, N+1)

	x := 0
	y := 0
	for i := 0; i < N; i++ {
		n := getInt()
		if 1 <= n && n <= N {
			a[n]++
			if a[n] != 1 {
				x = n
			}
		} else {
			y = n
		}
	}

	if y == 0 {
		for i := 1; i <= N; i++ {
			if a[i] == 0 {
				y = i
				break
			}
		}
	}
	if y == 0 {
		out("Correct")
	} else {
		out(x, y)
	}
}
