package main

import "fmt"

//Combination generator for int slice
func combinations(list []int, choose, buf int) (c chan []int) {
	c = make(chan []int, buf)
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
				for subComb := range combinations(list[i+1:], choose-1, buf) {
					c <- append([]int{list[i]}, subComb...)
				}
			}
		}
	}()
	return
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	buf := 5

	fmt.Println("Combinations")
	for comb := range combinations(a, 3, buf) {
		fmt.Println(comb)
	}
}
