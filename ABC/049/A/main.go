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

	c := getString()

	ans := "consonant"

	switch c[0] {
	case 'a':
		ans = "vowel"
	case 'i':
		ans = "vowel"
	case 'u':
		ans = "vowel"
	case 'e':
		ans = "vowel"
	case 'o':
		ans = "vowel"
	}
	out(ans)
}
