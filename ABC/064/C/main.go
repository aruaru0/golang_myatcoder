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

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	rate := make([]int, 9)
	for i := 0; i < N; i++ {
		a := getInt()
		switch {
		case a >= 1 && a <= 399:
			rate[0]++
		case a >= 400 && a <= 799:
			rate[1]++
		case a >= 800 && a <= 1199:
			rate[2]++
		case a >= 1200 && a <= 1599:
			rate[3]++
		case a >= 1600 && a <= 1999:
			rate[4]++
		case a >= 2000 && a <= 2399:
			rate[5]++
		case a >= 2400 && a <= 2799:
			rate[6]++
		case a >= 2800 && a <= 3199:
			rate[7]++
		case a >= 3200:
			rate[8]++
		}
	}

	ans := 0
	for i := 0; i < 8; i++ {
		if rate[i] != 0 {
			ans++
		}
	}
	min := ans
	if ans == 0 {
		min = 1
	}
	max := ans + rate[8]
	out(min, max)
}
