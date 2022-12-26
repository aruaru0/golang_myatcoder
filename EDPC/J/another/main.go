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

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 2)
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

var dp [310][310][310]float64

func loop(n1, n2, n3, N int) float64 {
	if dp[n1][n2][n3] >= 0 {
		return dp[n1][n2][n3]
	}
	if n1 == 0 && n2 == 0 && n3 == 0 {
		return 0.0
	}

	res := float64(0.0)
	if n1 > 0 {
		res += loop(n1-1, n2, n3, N) * float64(n1)
	}
	if n2 > 0 {
		res += loop(n1+1, n2-1, n3, N) * float64(n2)
	}
	if n3 > 0 {
		res += loop(n1, n2+1, n3-1, N) * float64(n3)
	}
	res += float64(N)
	res *= 1.0 / float64(n1+n2+n3)
	dp[n1][n2][n3] = res

	return res
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	n := [3]int{0, 0, 0}
	for i := 0; i < N; i++ {
		switch a[i] {
		case 1:
			n[0]++
		case 2:
			n[1]++
		case 3:
			n[2]++
		}
	}

	for i := 0; i < 310; i++ {
		for j := 0; j < 310; j++ {
			for k := 0; k < 310; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	ans := loop(n[0], n[1], n[2], N)
	out(ans)
}
