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

// MOD :
var MOD = 1000000007

func mpow(x, n int) int {
	ans := 1
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x % MOD
		}
		x = x * x % MOD
		n = n >> 1
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	n := (N + 1) / 2
	if N%2 == 0 {
		n++
	}
	A := make([]int, n)

	flag := true
	for i := 0; i < N; i++ {
		a := getInt()
		a = (a + 1) / 2

		A[a]++
		if a == 0 && A[a] >= 2 {
			flag = false
			break
		} else if A[a] > 2 {
			flag = false
			break
		}
	}

	if flag == false {
		out(0)
	} else {
		if N%2 == 1 {
			n = N - 1
		}
		n = N / 2
		out(mpow(2, n))
	}

}
