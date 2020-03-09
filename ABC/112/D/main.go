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

// 約数を列挙
func divisor(n int) []int {
	ret := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
			if i*i != n {
				ret = append(ret, n/i)
			}
		}
	}
	return ret
}

func a(M, N int, ch chan int) {
	ans := 1
	for i := N; i < M; i++ {
		if M%i == 0 {
			ans = M / i
			break
		}
	}
	ch <- ans
	//	close(ch)
}

func b(M, N int, ch chan int) {
	ans := 1
	m := M/N + 1
	for i := 1; i < m; i++ {
		if M%i == 0 && ans < i {
			ans = i
		}
	}
	ch <- ans
	//	close(ch)
}

func solve(M, N int) int {
	m := M / N
	d := divisor(M)
	max := -1
	for _, v := range d {
		if max < v && v <= m {
			max = v
		}
	}
	return max
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := getInt(), getInt()

	/*	c := make(chan int)
		go a(M, N, c)
		go b(M, N, c)

		res := <-c
	*/
	out(solve(M, N))
}
