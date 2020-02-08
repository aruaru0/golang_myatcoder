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

// Pfs :　素因数分解し、スライスを作成
func Pfs(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkDigit(n int) int {
	cnt := 1
	for {
		if n/10 == 0 {
			break
		}
		n /= 10
		cnt++
	}
	return cnt
}

func PfsMap(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2] = pfs[2] + 1
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i] = pfs[i] + 1
			n = n / i
		}
	}

	if n > 2 {
		pfs[n] = pfs[n] + 1
	}

	return pfs
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()

	ans := checkDigit(N)

	for a := 1; a*a <= N; a++ {
		b := N / a
		if N%a != 0 {
			continue
		}
		d := max(checkDigit(a), checkDigit(b))
		if ans > d {
			ans = d
		}
	}

	out(ans)
}
