package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

// Data :
type Data struct {
	X string
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
	return p[i].X < p[j].X
}

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	L := nextInt()
	s := make(Datas, N)

	for i := 0; i < N; i++ {
		s[i].X = nextString()
	}

	out(N, L, s)
	sort.Sort(s)
	out(s)

	ans := ""
	for i := 0; i < N; i++ {
		ans = ans + s[i].X
	}
	fmt.Println(ans)

}
