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
	sc.Buffer([]byte{}, 200100)

	N, mod := getInt(), getInt()
	S := getString()

	// 特殊処理
	if mod == 2 || mod == 5 {
		ans := 0
		for i := 0; i < N; i++ {
			val := int(S[i] - '0')
			if val%mod == 0 {
				ans += i + 1
			}
		}
		out(ans)
		return
	}

	m := make(map[int]int)
	cur := 0
	m[cur] = 1
	digit := 1
	ans := 0
	for i := N - 1; i >= 0; i-- {
		cur += digit * int(S[i]-'0')
		cur %= mod
		ans += m[cur]
		m[cur]++
		digit = digit * 10 % mod
	}
	out(ans)
}
