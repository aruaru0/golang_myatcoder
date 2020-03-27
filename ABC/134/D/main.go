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

func count(a []int, s, N int) int {
	cnt := 0
	for i := s + s; i <= N; i += s {
		cnt += a[i]
	}
	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N+1)
	b := make([]int, N+1)
	ans := 0
	for i := 1; i <= N; i++ {
		a[i] = getInt()
		if i > N/2 {
			b[i] = a[i]
			if a[i] == 1 {
				ans++
			}
		}
	}

	for i := N / 2; i > 0; i-- {
		c := count(b, i, N)
		if a[i] != c%2 {
			b[i] = 1
			ans++
		}
		//		out(i, c)
	}

	//out(a)
	//out(b)
	out(ans)
	if ans != 0 {
		for i := 1; i <= N; i++ {
			if b[i] == 1 {
				fmt.Print(i, " ")
			}
		}
		out()
	}
}
