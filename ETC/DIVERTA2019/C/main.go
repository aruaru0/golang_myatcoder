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
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	s := make([]string, N)

	for i := 0; i < N; i++ {
		s[i] = getString()
	}

	frontB := 0
	backA := 0
	pairAB := 0
	innerAB := 0
	for i := 0; i < N; i++ {
		if s[i][0] == 'B' && s[i][len(s[i])-1] == 'A' {
			pairAB++
		} else if s[i][0] == 'B' {
			frontB++
		} else if s[i][len(s[i])-1] == 'A' {
			backA++
		}
		for j := 1; j < len(s[i]); j++ {
			if s[i][j] == 'B' && s[i][j-1] == 'A' {
				innerAB++
			}
		}
	}

	ans := 0
	if pairAB != 0 {
		ans = pairAB - 1
		if frontB > 0 {
			ans++
			frontB--
		}
		if backA > 0 {
			ans++
			backA--
		}
	}
	ans += min(backA, frontB)
	ans += innerAB

	out(ans)

}
