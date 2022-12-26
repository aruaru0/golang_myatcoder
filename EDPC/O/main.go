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

func bitCount(bits int) int {

	bits = (bits & 0x55555555) + (bits >> 1 & 0x55555555)
	bits = (bits & 0x33333333) + (bits >> 2 & 0x33333333)
	bits = (bits & 0x0f0f0f0f) + (bits >> 4 & 0x0f0f0f0f)
	bits = (bits & 0x00ff00ff) + (bits >> 8 & 0x00ff00ff)
	return (bits & 0x0000ffff) + (bits >> 16 & 0x0000ffff)
}

const mod = 1000000007

func rec(N int, a [22][22]int) int {
	var dp [22][1 << 22]int
	dp[0][0] = 1
	for i := 0; i < N; i++ {
		for bit := 0; bit < 1<<uint(N); bit++ {
			// bitがちょうどi個立っている（ペアがちょうどi)のみ対象
			if bitCount(bit) == i {
				// ペアリングする女性を調べる
				for j := 0; j < N; j++ {
					// 女性が余っていて、かつペアリングできる場合
					if (bit&(1<<uint(j))) == 0 && a[i][j] != 0 {
						pos := bit | (1 << uint(j))
						// ペアリングした組み合わせをすべて足しこむ
						dp[i+1][pos] += dp[i][bit]
						dp[i+1][pos] %= mod
					}
				}
			}
		}
	}
	return dp[N][(1<<uint(N))-1]
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	var a [22][22]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = getInt()
		}
	}

	// 再帰で作りかけたが、こちらのほうが簡単っぽいので変更
	res := rec(N, a)
	out(res)
}
