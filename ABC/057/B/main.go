package main

import (
	"bufio"
	"fmt"
	"math"
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

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, M)
	d := make([]int, M)
	for i := 0; i < N; i++ {
		a[i], b[i] = getInt(), getInt()
	}
	for i := 0; i < M; i++ {
		c[i], d[i] = getInt(), getInt()
	}
	for i := 0; i < N; i++ {
		min := math.MaxInt64
		pos := 0
		for j := 0; j < M; j++ {
			diff := asub(a[i], c[j]) + asub(b[i], d[j])
			//out(a[i], b[i], c[i], d[i], "diff", diff, min, pos)
			if diff < min {
				min = diff
				pos = j
				//out("min", min, j)
			}
		}
		out(pos + 1)
	}
}
