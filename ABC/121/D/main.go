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

func f(a, b int) int {
	// 偶数、奇数の並びは xorすると１となる性質を使う
	n := b - a + 1
	ans := 0
	if a%2 == 1 { // 手前に偶数がない場合
		ans ^= a
		a++
		n--
	}
	d := n / 2
	if b%2 == 0 { // 次に偶数が続かないなら
		ans ^= (d % 2) ^ b
	} else {
		ans ^= d % 2
	}
	return ans
}

func f2(n int) int {
	ret := 0
	switch n % 4 {
	case 0:
		ret = n
	case 1:
		ret = 1
	case 2:
		ret = n ^ 1
	case 3:
		ret = 0
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	A, B := getInt(), getInt()

	out(f(A, B))
	// out(f2(A-1) ^ f2(B))
}
