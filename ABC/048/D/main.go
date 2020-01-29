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
	r := bufio.NewReaderSize(os.Stdin, 100000)
	s, _, _ := r.ReadLine()

	v := s[0]
	cnt := 0
	for i := 1; i < len(s)-1; i++ {
		if v == s[i+1] {
			v = s[i+1]
			i++
		} else {
			cnt++
		}
	}

	if cnt == 0 {
		out("Second")
	} else if cnt%2 == 0 {
		out("Second")
	} else {
		out("First")
	}
}
