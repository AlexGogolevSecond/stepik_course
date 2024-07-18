package main

import "fmt"

func main(){
	var a [3]int
	ta := fmt.Sprintf("%T", a)

	fmt.Println(a)  // [0 0 0]	
	fmt.Println(ta)	 // [3]int

	var b [5]int
	tb := fmt.Sprintf("%T", b)
	
	fmt.Println(b)  // [0 0 0 0 0]
	fmt.Println(tb)  // [5]int
}
