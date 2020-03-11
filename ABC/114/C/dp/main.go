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

func check(N int) int {
	ans := 0
	for i := 0; i < N; i++ {
		n := i
		var x, x3, x5, x7 int
		for n > 0 {
			if n%10 == 3 {
				x3 = 1
			}
			if n%10 == 5 {
				x5 = 1
			}
			if n%10 == 7 {
				x7 = 1
			}
			if n%10 != 3 && n%10 != 5 && n%10 != 7 {
				x = 1
			}
			n /= 10
		}
		if x3+x5+x7 == 3 && x == 0 {
			ans++
		}
	}
	return (ans)
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

	//  [digit][smaller][cond1][cond2][cond3]
	//  cond1 ... 3で含む
	//  cond2 ... 5を含む
	//  cond3 ... 7を含む
	//  cond4 ... 375以外を含む
	//  cond5 ... 0より大きい
	var dp [15][2][2][2][2][2][2]int

	dp[0][0][0][0][0][0][0] = 1
	for i := 0; i < nbits; i++ {
		for m := 0; m < 2; m++ {
			for j3 := 0; j3 < 2; j3++ {
				for j5 := 0; j5 < 2; j5++ {
					for j7 := 0; j7 < 2; j7++ {
						for z := 0; z < 2; z++ {
							for n357 := 0; n357 < 2; n357++ {
								for a := 0; a <= condAB(m != 0, 9, d[i]); a++ {
									mm := m
									if a < d[i] {
										mm = 1
									}
									jj3 := j3
									if a == 3 {
										jj3 = 1
									}
									jj5 := j5
									if a == 5 {
										jj5 = 1
									}
									jj7 := j7
									if a == 7 {
										jj7 = 1
									}
									nn357 := n357
									if z == 1 {
										if a != 3 && a != 5 && a != 7 {
											nn357 = 1
										}
									} else {
										if a != 0 && a != 3 && a != 5 && a != 7 {
											nn357 = 1
										}
									}
									zz := z
									if a > 0 {
										zz = 1
									}
									dp[i+1][mm][jj3][jj5][jj7][nn357][zz] += dp[i][m][j3][j5][j7][n357][z]
								}
							}
						}
					}
				}
			}
		}
	}

	ans := 0
	for m := 0; m < 2; m++ {
		for j3 := 0; j3 < 2; j3++ {
			for j5 := 0; j5 < 2; j5++ {
				for j7 := 0; j7 < 2; j7++ {
					for z := 0; z < 2; z++ {
						for n357 := 0; n357 < 2; n357++ {
							if j3 == 1 && j5 == 1 && j7 == 1 && n357 == 0 {
								ans += dp[nbits][m][j3][j5][j7][n357][z]
							}
						}
					}
				}
			}
		}
	}
	out(ans)
	//	out(check(N))
}
