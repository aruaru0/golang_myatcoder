package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func calc(x, y, z, i int) int {
	ret := 0
	if i%2 == 0 {
		ret += x
	} else {
		ret -= x
	}
	i /= 2
	if i%2 == 0 {
		ret += y
	} else {
		ret -= y
	}
	i /= 2
	if i%2 == 0 {
		ret += z
	} else {
		ret -= z
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	x := make([]int, N)
	y := make([]int, N)
	z := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i], z[i] = getInt(), getInt(), getInt()
	}

	ans := 0
	for i := 0; i < 8; i++ {
		a := make([]int, N)
		for j := 0; j < N; j++ {
			a[j] = calc(x[j], y[j], z[j], i)
		}
		sort.Ints(a)
		sum := 0
		for j := 0; j < M; j++ {
			sum += a[N-1-j]
		}
		ans = max(ans, sum)
	}
	out(ans)
}
