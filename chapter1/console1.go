package main

import "fmt"

func main(){
	var(
		name string
		age int

	)
	fmt.Print("Введите имя:")
	fmt.Scan(&name)
	fmt.Print("Введите возраст:")
	fmt.Scan(&age)

	fmt.Println(name, age)
}