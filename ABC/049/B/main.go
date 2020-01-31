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

	H, _ := getInt(), getInt()

	S := make([]string, H)
	for i := 0; i < H; i++ {
		S[i] = getString()
	}

	for i := 0; i < H; i++ {
		out(S[i])
		out(S[i])
	}
}
