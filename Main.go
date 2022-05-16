package main

import "fmt"

func main() {

	var i int = 3
	// var j string
	// j = i
	// var myName string = "hello"
	var j string = string(i)
	var k = 2.56
	fmt.Println(i, j)

	fmt.Printf("%T, %v", k, k)
	fmt.Print("\n")
}
