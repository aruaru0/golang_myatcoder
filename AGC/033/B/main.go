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

	H, W, N := getInt(), getInt(), getInt()
	sr, sc := getInt(), getInt()

	S, T := getString(), getString()

	// 先手はLのみ、後手はRのみ
	pos := sc
	for i := 0; i < N; i++ {
		//out(pos, string(S[i]), string(T[i]))
		if S[i] == 'L' {
			pos--
		}
		if pos == 0 {
			out("NO")
			return
		}
		if T[i] == 'R' {
			pos++
			if pos > W {
				pos = W
			}
		}
	}
	// 先手はRのみ、後手はLのみ
	pos = sc
	for i := 0; i < N; i++ {
		//out(pos, string(S[i]), string(T[i]))
		if S[i] == 'R' {
			pos++
		}
		if pos == W+1 {
			out("NO")
			return
		}
		if T[i] == 'L' {
			pos--
			if pos == 0 {
				pos = 1
			}
		}
	}
	// DU
	pos = sr
	for i := 0; i < N; i++ {
		if S[i] == 'U' {
			pos--
		}
		if pos == 0 {
			out("NO")
			return
		}
		if T[i] == 'D' {
			pos++
			if pos > H {
				pos = H
			}
		}
	}

	// UD
	pos = sr
	for i := 0; i < N; i++ {
		if S[i] == 'D' {
			pos++
		}
		if pos == H+1 {
			out("NO")
			return
		}
		if T[i] == 'U' {
			pos--
			if pos == 0 {
				pos = 1
			}
		}
	}

	out("YES")
}
