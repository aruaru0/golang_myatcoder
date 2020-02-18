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

func f(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	n := getInt()
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = getInt()
	}

	sign := -1
	ans := [2]int{0, 0}
	for k := 0; k < 2; k++ {
		if k == 0 {
			sign = -1
		} else {
			sign = 1
		}
		sum := 0 // sumの初期化位置のミスで結構苦労
		for i := 0; i < n; i++ {
			sum += a[i]
			if sign == 1 {
				if sum >= 0 {
					ans[k] += 1 + sum
					sum = -1
				}
			} else {
				if sum <= 0 {
					ans[k] += 1 - sum
					sum = 1
				}
			}
			sign = -sign
		}
	}

	fmt.Println(min(ans[0], ans[1]))
}
