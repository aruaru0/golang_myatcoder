package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 100000)

	s, _, _ := reader.ReadLine()
	str := string(s)

	fmt.Println(str)
}
