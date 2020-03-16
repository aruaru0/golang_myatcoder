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

	H, W := getInt(), getInt()
	var c [10][10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			c[i][j] = getInt()
		}
	}
	a := make(map[int]int)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			x := getInt()
			a[x]++
		}
	}

	for k := 0; k < 10; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				c[i][j] = min(c[i][j], c[i][k]+c[k][j])
			}
		}
	}

	ans := 0
	for i, n := range a {
		if i == -1 {
			continue
		}
		cost := c[i][1]
		ans += cost * n
	}

	out(ans)
}
