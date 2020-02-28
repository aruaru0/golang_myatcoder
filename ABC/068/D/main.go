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

func check(a, b uint64) {
	cnt := 0
	for a >= 2 || b >= 2 {
		if a > b {
			a = a - 2
			b++
		} else {
			b = b - 2
			a++
		}
		cnt++
		//out(cnt, ":", a, b)
	}
	out(cnt)
}

func calcN(n int) {
	N := 50
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}

	x := n / N
	m := n % N
	//	out(0, a)
	for i := 1; i <= m; i++ {
		min := 0
		for j := 0; j < N; j++ {
			if a[min] > a[j] {
				min = j
			}
			a[j]--
		}
		a[min] += N + 1
		/*	out(i, ":", a, (i-1)/N)
			if i%N == 0 {
				out("-----")
			}
		*/
	}
	for i := 0; i < N; i++ {
		a[i] += x
		fmt.Printf("%v ", a[i])
	}
	out()
	//	out(a)

}

func calc(n int) {
	N := 50
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}
	out(0, a)
	for i := 1; i <= n; i++ {
		min := 0
		for j := 0; j < N; j++ {
			if a[min] > a[j] {
				min = j
			}
			a[j]--
		}
		a[min] += N + 1
		out(i, ":", a, (i-1)/N)
		if i%N == 0 {
			out("-----")
		}

	}
}

func main() {
	sc.Split(bufio.ScanWords)
	K := getInt()
	out(50)
	calcN(K)
}
