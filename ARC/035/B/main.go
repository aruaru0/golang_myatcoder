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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	t := make([]int, N)
	m := make(map[int]int)
	for i := 0; i < N; i++ {
		t[i] = getInt()
		m[t[i]]++
	}
	sort.Ints(t)

	ans0 := 0
	cur := 0
	for i := 0; i < N; i++ {
		cur += t[i]
		ans0 += cur
	}
	out(ans0)

	ans1 := 1
	for _, v := range m {
		c := 1
		for i := 1; i <= v; i++ {
			c *= i
			c %= mod
		}
		ans1 *= c
		ans1 %= mod
	}
	out(ans1)
}
