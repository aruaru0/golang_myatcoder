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

// Data :
type Data struct {
	a, b int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].a < p[j].a
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	n := make(Datas, N)
	for i := 0; i < N; i++ {
		n[i] = Data{getInt(), getInt()}
	}

	sort.Sort(n)

	ans := 0
	cnt := 0
	for _, v := range n {
		if cnt+v.b >= K {
			ans = v.a
			break
		}
		cnt += v.b
	}

	out(ans)
}
