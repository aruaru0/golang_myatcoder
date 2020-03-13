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

// 約数を列挙
func divisor(n int) []int {
	ret := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i*i != n {
				ret = append(ret, n/i)
			}
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()

	d := divisor(N)
	sort.Ints(d)

	cnt := 0
	for _, v := range d {
		m := N/v - 1
		if m == 0 {
			continue
		}
		if N/m == N%m {
			cnt += m
		}
	}
	out(cnt)
}
