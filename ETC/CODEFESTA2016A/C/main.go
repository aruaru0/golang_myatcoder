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
	sc.Buffer([]byte{}, 100100)

	s := []byte(getString())
	K := getInt()

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			continue
		}
		v := s[i]
		d := int(1 + 'z' - v)
		if K >= d {
			K -= d
			s[i] = 'a'
		}
	}
	K %= 26
	l := len(s) - 1
	s[l] += byte(K)
	if s[l] > 'z' {
		s[l] = 'a'
	}
	out(string(s))

}
