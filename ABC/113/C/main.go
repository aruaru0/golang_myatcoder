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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type city struct {
	p, y int
}

// Datas :
type cities []city

func (p cities) Len() int {
	return len(p)
}

func (p cities) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p cities) Less(i, j int) bool {
	if p[i].p == p[j].p {
		return p[i].y < p[j].y
	}
	return p[i].p < p[j].p
}

func main() {
	sc.Split(bufio.ScanWords)

	_, M := getInt(), getInt()
	c := make(cities, M)
	for i := 0; i < M; i++ {
		c[i].p, c[i].y = getInt(), getInt()
	}
	s := make(cities, M)
	copy(s, c)
	sort.Sort(s)
	//out(s)

	x := make(map[city]int)
	cnt := 1
	p := s[0].p
	for i := 0; i < M; i++ {
		//out(s[i], p, cnt)
		if s[i].p == p {
			x[s[i]] = cnt
			cnt++
		} else {
			cnt = 1
			x[s[i]] = cnt
			p = s[i].p
			cnt++
		}
	}

	//out(x)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < M; i++ {
		p := c[i].p
		n := x[c[i]]
		fmt.Fprintf(w, "%06d%06d\n", p, n)
	}
}
