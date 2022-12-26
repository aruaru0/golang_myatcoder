package main

import "fmt"

// 入力をbitビット列に変換する
func makeBit(v, n, bit int) []int {
	bits := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		bits[i] = v % bit
		v /= bit
	}
	return bits
}

// 入力をビットに見立てて１が立っている部分で区分けする
// 区分けした数と、それぞれの区分を返す
func makeArea(v, n int) (int, []int) {
	area := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		area[i] = cnt
		if v%2 == 1 {
			cnt++
		}
		v /= 2
	}
	return area[len(area)-1] + 1, area
}

func main() {
	len := 4

	n := 1 << uint(len)

	for i := 0; i < n; i++ {
		r := makeBit(i, len, 2)
		fmt.Println(r)
	}

	for i := 0; i < n; i++ {
		c, r := makeArea(i, len)
		fmt.Println(r, c)
	}

	n = 4 * 4
	for i := 0; i < n; i++ {
		r := makeBit(i, 2, 4)
		fmt.Println(r)
	}

}
