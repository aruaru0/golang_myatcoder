package main

import (
	"bufio"
	"fmt"
	"os"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 100000)

	s, _, err := r.ReadLine()
	if err != nil {
		panic(err)
	}

	pr := -1
	prr := -1
	prrr := -1
	start := -1
	end := -1

	// AxAとAAのパターンを探す。見つかったら終了
	for i, v := range s {
		pr, prr, prrr = int(v), pr, prr
		if pr == prr {
			start = i
			end = i + 1
			break
		}
		if pr == prrr {
			start = i - 1
			end = i + 1
			break
		}
	}
	fmt.Println(start, end)
}
