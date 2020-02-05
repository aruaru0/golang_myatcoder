package main

import (
	"bufio"
	"fmt"
	"os"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

func main() {
	re := bufio.NewReaderSize(os.Stdin, 200000)

	s, _, _ := re.ReadLine()

	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			start = i
			break
		}
	}
	end := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'Z' {
			end = i + 1
			break
		}
	}
	out(end - start)
}
