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

	s := getString()

	ans := len(s)
	for i := 1; i < len(s); i++ {
		substr := s[0 : len(s)-i]
		l := len(substr)
		if l%2 != 0 {
			continue
		}
		if substr[0:l/2] == substr[l/2:l] {
			ans = l
			break
		}
		//		out(substr, substr[0:l/2], substr[l/2:l])
	}
	out(ans)
}
