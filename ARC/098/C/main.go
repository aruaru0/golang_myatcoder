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
	sc.Buffer([]byte{}, 1000000)

	N, s := getInt(), getString()

	e := make([]int, N)
	w := make([]int, N)
	sume := 0
	sumw := 0
	for i := 0; i < N; i++ {
		if s[i] == 'W' {
			sume++
		}
		if s[N-1-i] == 'E' {
			sumw++
		}
		e[i] = sume
		w[N-1-i] = sumw
	}

	m := N + 100000
	for i := 0; i < N; i++ {
		if e[i]+w[i] < m {
			m = e[i] + w[i]
		}
	}
	//out(e)
	//out(w)
	out(m - 1)
}
