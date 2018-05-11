package main

import "fmt"

func main() {
	//fmt.Printf("%d - %b -%#X", 42, 42, 42)
	for i := 1; i < 10; i++ {
		fmt.Printf("%d \t %b \t %#X  \t %q \n", i, i, i, i)
	}
	//fmt.Printf("%d \t %b \t %#X", 42, 42, 42)
}
