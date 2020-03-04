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

func solve(N, M int, S, C []int) {
	k := make([]int, N+1)
	f := make([]int, N+1)
	ng := false
	for i := 0; i < M; i++ {
		s, c := S[i], C[i]
		if N < s {
			ng = true
			break
		}
		if f[s] == 1 && k[s] != c {
			ng = true
			break
		}
		if s == 1 && c == 0 && N != 1 {
			ng = true
			break
		}
		f[s] = 1
		k[s] = c
	}

	if k[1] == 0 && N != 1 {
		k[1] = 1
	}

	if ng == true {
		out(-1)
	} else {
		for i := 1; i <= N; i++ {
			fmt.Print(k[i])
		}
		out()
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := getInt(), getInt()
	S := make([]int, M)
	C := make([]int, M)
	for i := 0; i < M; i++ {
		S[i], C[i] = getInt(), getInt()
	}
	solve(N, M, S, C)
}
