package main

import (
	"fmt"
	"sort"
)

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func main() {
	a := []int{2, 2, 5, 5, 9}

	sort.Ints(a)
	fmt.Println(a)

	ret := lowerBound(a, 2)
	fmt.Println(ret, a[ret:])
	ret = lowerBound(a, 4)
	fmt.Println(ret, a[ret:])
	ret = lowerBound(a, 5)
	fmt.Println(ret, a[ret:])
	ret = lowerBound(a, 7)
	fmt.Println(ret, a[ret:])
	ret = lowerBound(a, 100)
	fmt.Println(ret, a[ret:])

	ret = upperBound(a, 2)
	fmt.Println(ret, a[ret:])
	ret = upperBound(a, 4)
	fmt.Println(ret, a[ret:])
	ret = upperBound(a, 5)
	fmt.Println(ret, a[ret:])
	ret = upperBound(a, 7)
	fmt.Println(ret, a[ret:])
	ret = upperBound(a, 100)
	fmt.Println(ret, a[ret:])
}
