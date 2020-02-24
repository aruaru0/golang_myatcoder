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
	r := bufio.NewReaderSize(os.Stdin, 4096)

	var H, W int
	fmt.Scan(&H, &W)
	a := make([]string, H)
	for i := 0; i < H; i++ {
		a[i] = string(readLine(r))
	}
	for x := 0; x < W+2; x++ {
		fmt.Print("#")
	}
	fmt.Println()
	for y := 0; y < H; y++ {
		fmt.Printf("#%s#\n", a[y])
	}
	for x := 0; x < W+2; x++ {
		fmt.Print("#")
	}
	fmt.Println()
}
