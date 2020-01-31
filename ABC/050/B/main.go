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

type dk struct {
	P, X int
}

type dks []dk

func (p dks) Len() int {
	return len(p)
}

func (p dks) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p dks) Less(i, j int) bool {
	if p[i].P != p[j].P {
		return p[i].P < p[j].P
	}
	return p[i].X < p[j].X
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	T := make([]int, N)
	for i := 0; i < N; i++ {
		T[i] = getInt()
	}
	M := getInt()
	p := make([]int, M)
	m := make([]int, M)
	for i := 0; i < M; i++ {
		p[i], m[i] = getInt()-1, getInt()
	}

	for i := 0; i < M; i++ {
		ans := 0
		for j := 0; j < N; j++ {
			if p[i] == j {
				ans += m[i]
			} else {
				ans += T[j]
			}
		}
		out(ans)
	}
}
