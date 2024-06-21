package main

import (
	"fmt"
	p "playground/explore"
)

func main() {
	fmt.Println(p.AtoiBase("125", "0123456789"))
	fmt.Println(p.AtoiBase("1111101", "01"))
	fmt.Println(p.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(p.AtoiBase("uoi", "choumi"))
	fmt.Println(p.AtoiBase("bbbbbab", "-ab"))
}