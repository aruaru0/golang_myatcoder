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
	N := getInt()
	a := make(map[int]int)

	for i := 0; i < N; i++ {
		x := getInt()
		a[x]++
	}

	// out(a)

	if len(a) > 3 {
		out("No")
		return
	}

	if len(a) == 1 {
		for i := range a {
			if i == 0 {
				out("Yes")
				return
			}
		}
		out("No")
		return
	}

	if len(a) == 2 {
		x, ok := a[0]
		if !ok {
			out("No")
			return
		}
		if N/x == 3 && N%x == 0 {
			out("Yes")
			return
		}
	}

	if N%3 != 0 {
		out("No")
		return
	}
	N /= 3
	ans := -1
	for i, v := range a {
		if v != N {
			out("No")
			return
		}
		if ans == -1 {
			ans = i
		} else {
			ans = ans ^ i
		}
	}
	if ans == 0 {
		out("Yes")
	} else {
		out("No")
	}
}
