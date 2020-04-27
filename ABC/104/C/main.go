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

const inf = 1000000

func makeBit(n, N int) []bool {
	f := make([]bool, N)
	idx := 0
	for n > 0 {
		if n%2 == 1 {
			f[idx] = true
		}
		idx++
		n /= 2
	}
	return f
}

func main() {
	sc.Split(bufio.ScanWords)
	D, G := getInt(), getInt()/100
	p := make([]int, D)
	c := make([]int, D)
	s := make([]int, D)
	for i := 0; i < D; i++ {
		p[i] = getInt()
		c[i] = getInt() / 100
		s[i] = p[i]*(i+1) + c[i]
	}

	N := 1 << uint(D)
	ans := inf
	for i := 0; i < N; i++ {
		f := makeBit(i, D)
		// out("loop", i, f)
		g := G
		cnt := 0
		for j, v := range f {
			if v == true {
				g -= s[j]
				cnt += p[j]
			}
		}
		if g > 0 {
			// out("pass----")
			for j := D - 1; j >= 0; j-- {
				if f[j] == false {
					rest := (g + j) / (j + 1)
					rest = min(rest, p[j])
					g -= rest * (j + 1)
					cnt += rest
					// out(j+1, rest, g)
				}
				if g < 0 {
					break
				}
			}
		}
		ans = min(ans, cnt)
		// out("cnt=", cnt, "g = ", g)
	}

	out(ans)
}
