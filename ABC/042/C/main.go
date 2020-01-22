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

	N := getInt()
	K := getInt()
	D := make([]int, K)
	for i := 0; i < K; i++ {
		D[i] = getInt()
	}

	//	out(N, K, D)

	// 力技での解答
	ans := N
	for ans < N*10 {
		flag := false
	L1:
		for j := ans; j != 0; j /= 10 {
			v := j % 10
			for k := 0; k < K; k++ {
				if v == D[k] {
					flag = true
					break L1
				}
			}
		}
		if flag == false {
			break
		}
		//out(ans)
		ans++
	}
	out(ans)
}
