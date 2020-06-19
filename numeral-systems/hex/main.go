package main

import "fmt"

func writeHexidecimal(int, int, int) {
	fmt.Printf("%d \t %b \t %#X \n", 32, 32, 32)
}

func main() {
	writeHexidecimal(32, 32, 32)
}
