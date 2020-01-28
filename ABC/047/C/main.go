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
	//sc.Split(bufio.ScanWords)

	reader := bufio.NewReaderSize(os.Stdin, 100000)
	s, _, _ := reader.ReadLine()

	c := s[0]
	cnt := 0
	for i := 1; i < len(s); i++ {
		if c != s[i] {
			cnt++
		}
		c = s[i]
	}

	out(cnt)
}
