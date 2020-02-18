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

/*
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
	r := bufio.NewReaderSize(os.Stdin, 4096)
*/

func main() {
	sc.Split(bufio.ScanWords)

	A, B, C := getInt(), getInt(), getInt()

	ans := "No"
	if A == B && A != C {
		ans = "Yes"
	}
	if A == C && A != B {
		ans = "Yes"
	}
	if B == C && A != B {
		ans = "Yes"
	}
	out(ans)
}
