package main

import (
	"bufio"
	"fmt"
	"os"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func check(st []int, N int, s []byte) bool {

	for i := 2; i < N; i++ {
		a := st[i-2]
		b := st[i-1]
		if b == 0 { // 羊
			if s[i-1] == 'o' {
				st[i] = a
			} else {
				if a == 0 {
					st[i] = 1
				} else {
					st[i] = 0
				}
			}
		} else { // 狼
			if s[i-1] == 'x' {
				st[i] = a
			} else {
				if a == 0 {
					st[i] = 1
				} else {
					st[i] = 0
				}
			}
		}
	}

	ans0 := false
	// 先頭が正しいかチェック
	x := st[1]
	y := st[N-1]
	if st[0] == 0 {
		if s[0] == 'o' && x == y {
			ans0 = true
		} else if s[0] == 'x' && x != y {
			ans0 = true
		}
	}
	if st[0] == 1 {
		if s[0] == 'x' && x == y {
			ans0 = true
		} else if s[0] == 'o' && x != y {
			ans0 = true
		}
	}

	ans1 := false
	// 先頭が正しいかチェック
	x = st[N-2]
	y = st[0]
	if st[N-1] == 0 {
		if s[N-1] == 'o' && x == y {
			ans1 = true
		} else if s[N-1] == 'x' && x != y {
			ans1 = true
		}
	}
	if st[N-1] == 1 {
		if s[N-1] == 'x' && x == y {
			ans1 = true
		} else if s[N-1] == 'o' && x != y {
			ans1 = true
		}
	}

	return (ans0 && ans1)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 100000)

	var N int
	fmt.Scan(&N)
	s, _, _ := r.ReadLine()

	//out(N, string(s))

	st0 := make([]int, N)

	st0[0] = 0
	st0[1] = 0
	res := check(st0, N, s)
	if res == false {
		st0[0] = 1
		st0[1] = 0
		res = check(st0, N, s)
	}
	if res == false {
		st0[0] = 0
		st0[1] = 1
		res = check(st0, N, s)
	}
	if res == false {
		st0[0] = 1
		st0[1] = 1
		res = check(st0, N, s)
	}

	if res == false {
		out(-1)
	} else {
		w := bufio.NewWriter(os.Stdout)
		for i := 0; i < N; i++ {
			if st0[i] == 0 {
				fmt.Fprint(w, "S")
			} else {
				fmt.Fprint(w, "W")
			}
		}
		fmt.Fprint(w, "\n")
		w.Flush()
	}
}
