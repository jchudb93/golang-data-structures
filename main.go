package main

import "fmt"

func main() {
	var pc *int
	var c int
	c = 3
	pc = &c
	*pc = 2
	fmt.Println(c)
}
