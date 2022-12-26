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

func pmap(p [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] == 1 {
				fmt.Print("Y")
			} else {
				fmt.Print(".")
			}
		}
		out()
	}
	out("---------")
}

func main() {
	sc.Split(bufio.ScanWords)

	N, Q := getInt(), getInt()
	p := make([][]int, N)
	for i := 0; i < N; i++ {
		p[i] = make([]int, N)
	}

	for i := 0; i < Q; i++ {
		c := getInt()
		switch c {
		case 1:
			a, b := getInt()-1, getInt()-1
			p[a][b] = 1
		case 2:
			a := getInt() - 1
			for x := 0; x < N; x++ {
				if p[x][a] == 1 {
					p[a][x] = 1
				}
			}
		case 3:
			a := getInt() - 1
			tmp := make([]int, N)
			copy(tmp, p[a])
			for y, v := range tmp {
				if v == 1 {
					for x := 0; x < N; x++ {
						if x != a && p[y][x] == 1 {
							p[a][x] = 1
						}
					}
				}
			}
		}

		//pmap(p, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if p[i][j] == 1 {
				fmt.Print("Y")
			} else {
				fmt.Print("N")
			}
		}
		out()
	}
}
