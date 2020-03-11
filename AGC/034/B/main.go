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
	sc.Buffer([]byte{}, 200200)

	s := []byte(getString())

	cnt := 0
	pos := 0
	bcc := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			out("A")
			cnt++
		} else if i < len(s)-1 && s[i] == 'B' && s[i+1] == 'C' {
			out("BC", cnt)
			i++
			pos += cnt - bcc
			cnt++
			bcc++
		} else {
			if bcc != 0 {
				sum += pos
			}
			out("sum", sum, "pos", pos, "cnt", cnt, bcc)
			cnt = 0
			bcc = 0
			pos = 0
		}
	}
	out("sum", sum, "pos", pos, "cnt", cnt)
	if bcc != 0 {
		sum += pos
	}
	fmt.Println(sum)
}
