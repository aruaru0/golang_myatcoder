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

	var dp [301][301][301]float64
	for i := 0; i < 301; i++ {
		for j := 0; j < 301; j++ {
			for k := 0; k < 301; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	dp[0][0][0] = 0.0
	for k := 0; k <= n[2]; k++ {
		for j := 0; j <= n[1]+n[2]-k; j++ {
			for i := 0; i <= n[0]+n[1]+n[2]-k-j; i++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				sum := 0.0
				if i > 0 {
					sum += dp[i-1][j][k] * float64(i)
				}
				if j > 0 {
					sum += dp[i+1][j-1][k] * float64(j)
				}
				if k > 0 {
					sum += dp[i][j+1][k-1] * float64(k)
				}
				sum += float64(N)
				sum /= float64(i + j + k)
				dp[i][j][k] = sum
			}
		}
	}

	out(dp[n[0]][n[1]][n[2]])
}
