package main

import (
	"fmt"
	"sort"
)

// Data : 
type Data struct {
	X int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].X < p[j].X
}

/*
func main() {
	//	fmt.Println("start")
	var N int
	sc.Split(bufio.ScanWords)
	N = nextInt()
	d := make(Datas, N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		d[i].X = nextInt()
		a[i] = d[i].X
	}

	sort.Sort(d)
	sort.Ints(a)
	fmt.Println(N, d, a)
}
}


