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

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type xy struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	x := make([]int, N)
	y := make([]int, N)
	pt := make(map[xy]int)
	for i := 0; i < N; i++ {
		x[i], y[i] = getInt(), getInt()
		pt[xy{x[i], y[i]}] = 1
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			dx := x[i] - x[j]
			dy := y[i] - y[j]
			cnt := 0
			// dx,dy（設問のp,q）の距離にあるものをカウント
			for k := 0; k < N; k++ {
				px := dx + x[k]
				py := dy + y[k]
				_, ok := pt[xy{px, py}]
				if ok {
					cnt++
				}
			}
			if ans < cnt {
				ans = cnt
			}
		}
	}
	out(N - ans)
}
