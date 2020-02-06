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
	N, M := getInt(), getInt()
	a := make([]string, N)
	b := make([]string, M)
	for i := 0; i < N; i++ {
		a[i] = getString()
	}
	for i := 0; i < M; i++ {
		b[i] = getString()
	}
	//out(N, M, a, b)

	ans := "No"
	for i := 0; i <= N-M; i++ {
		for j := 0; j <= N-M; j++ {
			ok := true
		L0:
			for k := 0; k < M; k++ {
				for l := 0; l < M; l++ {
					//out(i, j, k, l, a[i+k][j+l], b[k][l])
					if a[i+k][j+l] != b[k][l] {
						//out("!=")
						ok = false
						break L0
					}
				}
			}
			if ok == true {
				ans = "Yes"
			}
		}
	}

	out(ans)
}
