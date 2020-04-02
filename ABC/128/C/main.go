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

func bitCount(bits int) int {

	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()
	s := make([]int, M)
	for i := 0; i < M; i++ {
		k := getInt()
		for j := 0; j < k; j++ {
			b := getInt() - 1
			s[i] |= 1 << uint(b)
		}
	}
	p := make([]int, M)
	for i := 0; i < M; i++ {
		p[i] = getInt()
	}

	// for i := 0; i < M; i++ {
	// 	fmt.Printf("%d %b\n", i, s[i])
	// }

	bit := 1 << uint(N)
	ans := 0
	for i := 0; i < bit; i++ {
		ok := true
		for j := 0; j < M; j++ {
			//fmt.Printf("%b %d\n", i, j)
			x := i & s[j]
			y := bitCount(x)
			if y%2 != p[j] {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	out(ans)
}
