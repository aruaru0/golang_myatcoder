package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(strings.Count(s, "(^^*)"), strings.Count(s, "(*^^)"))
}
