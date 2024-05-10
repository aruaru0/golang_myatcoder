package main

import "fmt"

func main() {
	var h, w int
	fmt.Scanf("%d %d", &h, &w)

	if h*h*20 > w*10000 {
		fmt.Println("A")
	} else if h*h*25 <= w*10000 {
		fmt.Println("C")
	} else {
		fmt.Println("B")
	}
}
