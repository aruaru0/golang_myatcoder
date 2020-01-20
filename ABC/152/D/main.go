package main

import "fmt"

// Out :
func Out(x ...interface{}) {
	//	fmt.Println(x...)
}

// GetInt :
func GetInt() int {
	var n int
	fmt.Scanf("%d", &n)
	return n
}

func main() {

	N := GetInt()

	Out(N)

	var c [10][10]int

	for i := 1; i <= N; i++ {
		lsb := i % 10
		msb := i % 10
		for v := i; v != 0; {
			msb = v % 10
			v /= 10
		}
		c[msb][lsb]++
		//	Out(i, msb, lsb)
	}

	//Out(c)
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			sum += c[i][j] * c[j][i]
		}
	}
	fmt.Println(sum)
}
