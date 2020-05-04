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

// min, max, asub, absなど基本関数
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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

type area struct {
	x, y int
}

func check(e []area, M, N int) {
	m := make([][]bool, N)
	n := make(map[int]int)
	for i := 0; i < N; i++ {
		m[i] = make([]bool, N)
		n[i] = i
	}
	out(e)
	out(N)
	for i := 0; i < N; i++ {
		out("----", i)
		for j := 0; j < M; j++ {
			a := n[e[j].x]
			b := n[e[j].y]
			out(a, b)
			if a == b {
				out("error")
				return
			}
			if m[a][b] == false && m[b][a] == false {
				m[a][b] = true
				m[b][a] = true
			} else {
				out("error")
				return
			}
		}
		for j := 0; j < N; j++ {
			n[j] = (i + j + 1) % N
		}
	}
	out("OK")
}

func main() {
	sc.Split(bufio.ScanWords)
	_, M := getInt(), getInt()

	e := make([]area, M)
	idx := 0
	// for i := 0; i < M; i++ {
	// 	a := i + 1
	// 	b := 2*M - i
	// 	e[i] = area{a - 1, b - 1}
	// 	out(a, b, abs(a-b), N-abs(a-b))
	// }
	// check(e, M, N)

	for i := 0; i < M; i++ {
		a := i + 1
		b := M - i
		if a < b {
			e[idx] = area{a - 1, b - 1}
			idx++
			out(a, b)
		} else {
			break
		}
	}

	for i := 0; i < M; i++ {
		a := M + i + 1
		b := 2*M - i + 1
		if a < b {
			e[idx] = area{a - 1, b - 1}
			idx++
			out(a, b)
		} else {
			break
		}
	}
	// check(e, M, N)

}
