package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)

	for i := 1; i < 101; i++ {
		if i*i*i == a {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
	return
}
