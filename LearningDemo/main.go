package main

import "C"

func main() {}

//export handler
func handler(a, b C.int) C.int {
	return a + b
}
