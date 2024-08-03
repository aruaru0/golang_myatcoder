package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 49
const MOD int = 998244353

type Matrix [][]int

func mul(A, B Matrix) Matrix {
	C := make(Matrix, N)
	for i := 0; i < N; i++ {
		C[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if A[j][i] != 0 {
				for k := 0; k < N; k++ {
					C[j][k] = (C[j][k] + A[j][i]*B[i][k]) % MOD
				}
			}
		}
	}
	return C
}

func main() {
	initMatrix := make(Matrix, N)
	for i := 0; i < N; i++ {
		initMatrix[i] = make([]int, N)
	}

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			x := 52
			if i != 6 {
				initMatrix[i*7+j][i*7+j+7]++
				x--
			}
			if j != 6 {
				initMatrix[i*7+j][i*7+j+1]++
				x--
			}
			initMatrix[i*7+j][i*7+j] += x
		}
	}

	pw := make([]Matrix, 31)
	pw[0] = initMatrix
	for i := 1; i < 31; i++ {
		pw[i] = mul(pw[i-1], pw[i-1])
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var t int
	fmt.Sscan(scanner.Text(), &t)

	for t > 0 {
		t--
		scanner.Scan()
		var n int
		fmt.Sscan(scanner.Text(), &n)

		v := make(Matrix, N)
		for i := 0; i < N; i++ {
			v[i] = make([]int, N)
		}
		v[0][0] = 1

		for b := 30; b >= 0; b-- {
			if (n>>b)&1 == 1 {
				v = mul(v, pw[b])
			}
		}

		fmt.Println(v[0][48])
	}
}
