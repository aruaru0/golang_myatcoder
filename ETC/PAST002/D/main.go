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

func main() {
	sc.Split(bufio.ScanWords)
	s := getString()
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := 0; j < 3; j++ {
			if i+j >= len(s) {
				continue
			}
			x := s[i : i+j+1]
			n := 1 << uint(len(x))
			for k := 0; k < n; k++ {
				tmp := []byte(x)
				b := k
				idx := 0
				for b > 0 {
					if b%2 == 1 {
						tmp[idx] = '.'
					}
					b /= 2
					idx++
				}
				m[string(tmp)]++
				// out(string(tmp))
			}
			// out(x)
		}
	}
	// out(m)
	out(len(m))
}
