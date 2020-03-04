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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	s := make([][]byte, 2)
	s[0] = []byte(getString())
	s[1] = []byte(getString())

	sum := 1
	i := 0
	prev := 0
	if s[0][0] != s[1][0] { // 横
		sum = 3 * 2
		i = 2
		prev = 1
	} else { // 縦
		sum = 3
		i = 1
		prev = 2
	}
	for i < N {
		if s[0][i] != s[1][i] { // 上下に水平
			if prev == 1 {
				sum = (sum * 3) % mod
			} else {
				sum = (sum * 2) % mod
			}
			prev = 1
			i += 2
		} else {
			if prev == 2 {
				sum = (sum * 2) % mod
			}
			prev = 2
			i++
		}
	}
	out(sum)
}
