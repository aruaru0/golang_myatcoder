package main

import (
	"bufio"
	"fmt"
	"math"
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

func f(b, n int) int {
	if b > n {
		return n
	}
	return f(b, n/b) + n%b
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	S := getInt()

	ans := -1
	if N == S {
		ans = N + 1
		fmt.Println(ans)
		return
	}
	m := int(math.Sqrt(float64(N))) + 1
	for b := 2; b < m; b++ {
		if f(b, N) == S {
			ans = b
			break
		}
	}

	div := N - S
	if (ans == -1) && (div > 0) {
		for p := 1; p < m; p++ {
			if div%p != 0 {
				continue
			}
			b := div/p + 1
			if b < m {
				continue
			}
			q := N % b
			s := p + q
			if (s == S) && ((ans > b) || (ans == -1)) {
				ans = b
			}
		}
	}

	fmt.Println(ans)
}
