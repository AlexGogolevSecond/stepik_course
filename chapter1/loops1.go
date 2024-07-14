package main

import "fmt"


func main(){
	sum := 0
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("=============")

	var j = 1
	for ; j < 10; j++{
		//fmt.Println(j)
	    fmt.Println(j * j)
	}

	fmt.Println("==============")
	//аналог цикла while !!!
	var k = 1
	for k < 10 {
		fmt.Println(k * k)
		k++
	}


}
