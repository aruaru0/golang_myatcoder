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

func getCard(i int, s string) (int, int, int) {
	mark := 0
	switch s[i] {
	case 'S':
		mark = 0
	case 'H':
		mark = 1
	case 'D':
		mark = 2
	case 'C':
		mark = 3
	}
	i++
	num := 0
	switch s[i] {
	case '1':
		num = 10
		i++
	case '2':
		num = 2
	case '3':
		num = 3
	case '4':
		num = 4
	case '5':
		num = 5
	case '6':
		num = 6
	case '7':
		num = 7
	case '8':
		num = 8
	case '9':
		num = 9
	case 'J':
		num = 11
	case 'Q':
		num = 12
	case 'K':
		num = 13
	case 'A':
		num = 14
	}
	i++
	return mark, num, i
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	s := getString()

	// どの種類のロイヤルストレートフラッシュができるか検査
	var card [4][5]int
	var flash [4]int

	sel := -1
	for i := 0; i < len(s); {
		mark, num, next := getCard(i, s)
		//out(mark, num)
		if num >= 10 && card[mark][num-10] == 0 {
			card[mark][num-10] = 1
			flash[mark]++
			if flash[mark] == 5 {
				sel = mark
				break
			}
		}
		i = next
	}

	//out(sel)
	var get [5]int
	cnt := 0
	outcnt := 0
	for i := 0; i < len(s); {
		mark, num, next := getCard(i, s)
		if num >= 10 && mark == sel && get[num-10] == 0 {
			get[num-10] = 1
			cnt++
		} else {
			fmt.Print(s[i:next])
			outcnt++
		}
		if cnt == 5 {
			break
		}
		i = next
	}
	if outcnt == 0 {
		out(0)
	} else {
		out()
	}
}
