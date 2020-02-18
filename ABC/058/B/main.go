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

	O, E := getString(), getString()

	n := len(O) + len(E)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("%c", O[i/2])
		} else {
			fmt.Printf("%c", E[i/2])
		}
	}
	fmt.Println("")

}
