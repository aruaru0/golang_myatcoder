package main

import (
	"bufio"
	"fmt"
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

	var H, W int

	fmt.Scan(&H)
	fmt.Scan(&W)
	s := make([]string, H)

	for i := 0; i < H; i++ {
		s[i] = string(readLine(r))
	}

	//out(H, W, s)
	ans := "Yes"
L:
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] == '#' {
				flg := 0
				if x != 0 {
					if s[y][x-1] == '#' {
						flg = 1
					}
				}
				if x != W-1 {
					if s[y][x+1] == '#' {
						flg = 1
					}
				}
				if y != 0 {
					if s[y-1][x] == '#' {
						flg = 1
					}
				}
				if y != H-1 {
					if s[y+1][x] == '#' {
						flg = 1
					}

				}
				if flg == 0 {
					ans = "No"
					break L
				}
			}
		}
	}
	out(ans)
}
