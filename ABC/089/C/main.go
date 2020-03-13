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

//Combination generator for int slice
func combinations(list []int, choose int) (c chan []int) {
	c = make(chan []int, 1)
	go func() {
		defer close(c)
		switch {
		case choose == 0:
			c <- []int{}
		case choose == len(list):
			c <- list
		case len(list) < choose:
			return
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range combinations(list[i+1:], choose-1) {
					c <- append([]int{list[i]}, subComb...)
				}
			}
		}
	}()
	return
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	n := make([]int, 5)
	for i := 0; i < N; i++ {
		s := getString()
		switch s[0] {
		case 'M':
			n[0]++
		case 'A':
			n[1]++
		case 'R':
			n[2]++
		case 'C':
			n[3]++
		case 'H':
			n[4]++
		}
	}

	a := []int{0, 1, 2, 3, 4}

	ans := 0
	for comb := range combinations(a, 3) {
		x := 1
		for _, v := range comb {
			x *= n[v]
		}
		ans += x
	}
	out(ans)
}
