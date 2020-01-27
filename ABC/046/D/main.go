package main

import (
	"bufio"
	"fmt"
	"os"
)

func out(x ...interface{}) {
	//	fmt.Println(x...)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100000)

	s, _, _ := reader.ReadLine()
	out(s)

	g := 0
	p := 0
	cnt := 0
	for i := 0; i < len(s); i++ {
		out(string(s[i]))
		if s[i] == 'g' {
			if g > p {
				p++
				cnt++
				out("P")
			} else {
				g++
				out("G")
			}
		} else {
			if g > p {
				p++
				out("P")
			} else {
				g++
				cnt--
				out("G")
			}
		}
	}
	fmt.Println(cnt)
}
