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

	Sa := getString()
	Sb := getString()
	Sc := getString()

	win := ""
	next := byte('a')
L:
	for {
		out("now", string(next), ":", string(Sa), string(Sb), string(Sc))
		switch next {
		case 'a':
			if len(Sa) == 0 {
				win = "A"
				break L
			}
			next = Sa[0]
			Sa = Sa[1:]
			//			out("A", Sa, "next", string(next))
		case 'b':
			if len(Sb) == 0 {
				win = "B"
				break L
			}
			next = Sb[0]
			Sb = Sb[1:]
			out("B", Sb, "next", string(next))
		case 'c':
			if len(Sc) == 0 {
				win = "C"
				break L
			}
			next = Sc[0]
			Sc = Sc[1:]
			//			out("C", Sc, "next", string(next))
		}
	}

	fmt.Println(win)
}
