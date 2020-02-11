package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
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

	N, K := getInt(), getInt()
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getInt()
	}
	q := p[0:K]

	ans := float64(0.0)
	for i := 0; i < K; i++ {
		ans += (1.0 + float64(p[i])) / 2
		out(ans)
	}
	max := ans
	for i := K; i < N; i++ {
		ans = ans + (1.0+float64(p[i]))/2 - (1.0+float64(p[i-K]))/2
		if ans > max {
			max = ans
		}
	}
	out(N, K, p, q, max)
	fmt.Printf("%f\n", max)
}
