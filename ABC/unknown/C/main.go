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

// Datas :
type Datas []byte

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	m := make(map[string]int)

	for i := 0; i < N; i++ {
		s := Datas(getString())
		sort.Sort(s)
		m[string(s)]++
	}

	ans := 0
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	out(ans)
}
