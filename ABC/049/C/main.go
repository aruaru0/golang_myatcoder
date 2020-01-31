package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 100000)

	s, _, _ := r.ReadLine()

	str := string(s)

	ans := "YES"
	for i := 0; i < len(str); {
		if strings.HasPrefix(str[i:], "dreamerase") {
			i += 5
		} else if strings.HasPrefix(str[i:], "dreamer") {
			i += 7
		} else if strings.HasPrefix(str[i:], "eraser") {
			i += 6
		} else if strings.HasPrefix(str[i:], "dream") {
			i += 5
		} else if strings.HasPrefix(str[i:], "erase") {
			i += 5
		} else {
			ans = "NO"
			break
		}
	}
	out(ans)
}
