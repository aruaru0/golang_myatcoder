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

const MAX = 55555

var pN [MAX]int

func initPrime() {
	for i := 2; i < MAX/2; i++ {
		for j := i * 2; j < MAX; j += i {
			if pN[j] == 0 {
				pN[j] = i
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	initPrime()

	N := getInt()

	cnt := 0
	for i := 2; i < MAX; i++ {
		// 素数で、かつ、1引いたら５で割れる（５つ足すと５で割れる）
		if pN[i] == 0 && (i-1)%5 == 0 {
			fmt.Print(i, " ")
			cnt++
		}
		if cnt == N {
			break
		}
	}
	fmt.Println("")
}
