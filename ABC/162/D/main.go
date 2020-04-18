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

func solve(N int, s []byte) {
	ans := 0
	c := make([][3]int, N)
	r := 0
	g := 0
	b := 0
	for i := N - 1; i >= 0; i-- {
		switch s[i] {
		case 'R':
			r++
		case 'G':
			g++
		case 'B':
			b++
		}
		c[i][0] = r
		c[i][1] = g
		c[i][2] = b
	}
	for i := 0; i < N; i++ {
		switch s[i] {
		case 'R':
			// out("R", c[i][1], c[i][2])
			ans += c[i][1] * c[i][2]
		case 'G':
			// out("G", c[i][0], c[i][2])
			ans += c[i][0] * c[i][2]
		case 'B':
			// out("B", c[i][0], c[i][1])
			ans += c[i][0] * c[i][1]
		}
	}
	// out(ans)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			k := j + j - i
			if j < N && k < N {
				if s[i] != s[j] && s[j] != s[k] && s[i] != s[k] {
					ans--
				}
			}
		}
	}

	out(ans)
}

func solve2(N int, s []byte) {
	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if s[i] == s[j] {
				continue
			}
			for k := j + 1; k < N; k++ {
				if s[i] == s[k] || s[j] == s[k] {
					continue
				}
				if j-i != k-j {
					ans++
				}
			}
		}
	}
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	s := getString()

	// N := 4000
	// s := make([]byte, N)
	// for i := 0; i < N; i++ {
	// 	switch i % 3 {
	// 	case 0:
	// 		s[i] = 'R'
	// 	case 1:
	// 		s[i] = 'G'
	// 	case 2:
	// 		s[i] = 'B'
	// 	}
	// }

	solve(N, []byte(s))

}
