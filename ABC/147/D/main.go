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

// min, max, asub, absなど基本関数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	b := make([][]int, 61)
	for i := 0; i < 61; i++ {
		b[i] = make([]int, N+1)
		for j := 0; j < N; j++ {
			b[i][j+1] = b[i][j] + (a[j]>>uint(i))%2
		}
	}

	ans := 0
	for k := 0; k < 61; k++ {
		cnt := 0
		for j := 0; j < N-1; j++ {
			s := (a[j] >> uint(k)) % 2
			n := N - (j + 1)
			n1 := b[k][N] - b[k][j+1]
			n0 := n - n1
			if s == 0 {
				cnt += n1
			} else {
				cnt += n0
			}
			// out(s, n, n1, n0, cnt)
		}
		for i := 0; i < k; i++ {
			cnt *= 2
			cnt %= mod
		}
		// out(cnt)
		ans += cnt
		ans %= mod
	}
	out(ans)
}
