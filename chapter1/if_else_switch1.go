package main

import "fmt"

func main(){

	var a = 5
	var b = 4

	if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a == b")
	}

	var i int
	fmt.Scan(&i)
	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
	}
}
