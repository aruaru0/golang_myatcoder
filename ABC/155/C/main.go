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

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	s := make([]string, N)
	m := make(map[string]int)
	max := 0
	for i := 0; i < N; i++ {
		s[i] = getString()
		m[s[i]]++
		if max < m[s[i]] {
			max = m[s[i]]
		}
	}
	sel := make([]string, 0)
	for i, v := range m {
		if v == max {
			sel = append(sel, i)
		}
	}

	sort.Strings(sel)
	for _, v := range sel {
		fmt.Println(v)
	}
}
