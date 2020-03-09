package main

import (
	"bufio"
	"fmt"
	"os"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

func readLine(r *bufio.Reader) []byte {
	buf := make([]byte, 0, 1024)
	for {
		l, p, e := r.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return buf
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 200100)

	s := readLine(r)

	S := 0
	T := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'S' {
			if T == 0 {
				S++
			} else {
				T--
			}
		} else {
			T++
		}
	}
	out(T + S)
}
