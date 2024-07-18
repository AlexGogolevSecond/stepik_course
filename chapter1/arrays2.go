package main

import "fmt"

func main(){
	var a = [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	d := [3]int{1: 12}
	
	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12 0]
}
