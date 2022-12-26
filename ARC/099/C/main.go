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

const inf = 1001001001

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, K := getInt(), getInt()
	a := make([]int, N+1)
	m := inf
	for i := 0; i < N; i++ {
		a[i] = getInt()
		m = min(m, a[i])
	}
	// 1...Nの並び替えでなくてもOKなのでオーバーキルかも
	ans := 0
	flg := false
	for i := 0; i < N-1; {
		if i+1 < N && flg == true && a[i+1] == m {
			i++
		}
		ans++
		for k := 0; k < K; k++ {
			if i+k > N {
				break
			}
			if a[i+k] == m {
				flg = true
			}
		}
		i += (K - 1)
		// out(i)
	}
	out(ans)

}
