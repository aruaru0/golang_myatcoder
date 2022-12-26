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

func condAB(c bool, a, b int) int {
	if c {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()

	n := N
	nbits := 0
	d := make([]int, 0)
	for n > 0 {
		d = append([]int{n % 10}, d...)
		n /= 10
		nbits++
	}

	//  [digit][smaller][cond1]
	//  cond1 ... 3である
	var dp [10][2][2]int

	dp[0][0][0] = 1
	for i := 0; i < nbits; i++ {
		for m := 0; m < 2; m++ {
			for j := 0; j < 2; j++ {
				for a := 0; a <= condAB(m != 0, 9, d[i]); a++ {
					mm := m
					if a < d[i] {
						mm = 1
					}
					// cond1
					jj := 0
					if a == 3 {
						jj = 1
					}
					dp[i+1][mm][jj] += dp[i][m][j]
				}
			}
		}
	}

	for i := 0; i <= nbits; i++ {
		out(dp[i])
	}
}
