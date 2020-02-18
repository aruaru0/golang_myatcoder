package main

import (
	"bufio"
	"fmt"
	"math"
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
	r := bufio.NewReaderSize(os.Stdin, 4096)

	var n int

	fmt.Scan(&n)
	var d [50][26]int

	for i := 0; i < n; i++ {
		str := readLine(r)
		for _, v := range str {
			d[i][v-'a']++
		}
	}
	var res [26]int
	for j := 0; j < 26; j++ {
		min := math.MaxInt32
		for i := 0; i < n; i++ {
			if min > d[i][j] {
				min = d[i][j]
			}
		}
		res[j] = min
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < res[i]; j++ {
			fmt.Printf("%c", i+'a')
		}
	}
	fmt.Println("")

}
