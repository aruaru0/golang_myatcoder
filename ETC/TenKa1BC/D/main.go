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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N, K := getInt(), getInt()
	bit := make([]int, 31)

	// Kのビット構成を調査
	n := 0
	for k := K; k > 0; {
		bit[n] = k % 2
		n++
		k >>= 1
	}

	// K以下で1のビットが０になるパターンを列挙
	pat := make([]int, 32)
	mask := 0
	pat[0] = K
	idx := 1
	for i := 30; i >= 0; i-- {
		if bit[i] == 1 {
			pat[idx] = mask | (1<<uint(i) - 1)
			idx++
			mask |= 1 << uint(i)
		}
	}

	// すべてのパターンでチェック O(N)
	x := make([]int, idx)
	sum := make([]int, idx)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		for j := 0; j < idx; j++ {
			if a|pat[j] == pat[j] {
				sum[j] += b
				x[j] |= a
			}
		}
	}

	// 最大を調べる
	ans := 0
	for i := 0; i < idx; i++ {
		if sum[i] > ans {
			ans = sum[i]
		}
	}
	out(ans)
}
