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
	m := make(map[int]int)
	sum := 0
	m[sum]++
	for i := 0; i < N; i++ {
		a := getInt()
		sum += a
		m[sum]++
	}

	ans := 0
	for _, v := range m {
		if v >= 2 {
			if v%2 == 0 {
				ans += v / 2 * (v - 1)
			} else {
				ans += v * ((v - 1) / 2)
			}
		}
	}

	out(ans)
}
