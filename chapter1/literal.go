//литералы и строки


package main

import "fmt"

func main(){
	var s1 string
	s1 = `Sammy says, "Hello!"`
	fmt.Println(s1)  // Sammy says, "Hello!"

	var s2 string
	s2 = "Sammy likes the `fmt` package for formatting strings.."
	fmt.Println(s2)  // Sammy likes the `fmt` package for formatting strings..

	fmt.Println("Sammy says, \"Hello!\"")  // Sammy says, "Hello!"
}
