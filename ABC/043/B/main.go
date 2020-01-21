package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
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
	var ans string

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			ans += "0"
		} else if s[i] == '1' {
			ans += "1"
		} else if len(ans) != 0 {
			ans = ans[0 : len(ans)-1]
		}
	}
	fmt.Println(ans)

}
