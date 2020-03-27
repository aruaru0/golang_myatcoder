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

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			if i == j {
				continue
			}
			for k := j; k < N; k++ {
				if i == k || j == k {
					continue
				}
				a := x[j] - x[i]
				b := y[j] - y[i]
				c := x[k] - x[i]
				d := y[k] - y[i]
				n := abs(a*d - b*c)
				if n != 0 && n%2 == 0 {
					ans++
				}
			}
		}
	}
	out(ans)
}
