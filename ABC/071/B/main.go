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

	S := getString()
	m := make([]int, 26)
	for _, v := range S {
		m[v-'a'] = 1
	}
	ans := "None"
	for i := 0; i < 26; i++ {
		if m[i] == 0 {
			ans = string('a' + i)
			break
		}
	}
	out(ans)
}
