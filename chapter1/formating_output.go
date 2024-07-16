package main

import "fmt"

func main(){
	var a rune = 'Ы'  // !!!???? так можно???
	fmt.Println(a)  // тут выведется 1067
	fmt.Printf("%q", a)  // 'Ы'

	// вывод: 'Ы'
	var s string = "abcde"
	fmt.Printf("%q", s)  // выведет "abcde"

}
