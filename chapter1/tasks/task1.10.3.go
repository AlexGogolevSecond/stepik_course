/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:

5
38 24 800 8 16
Sample Output:

40
*/
package main

import "fmt"


func main(){
	var (
		c int
		i int
		number int
		sum int
	)
	
	fmt.Scan(&c)
	// fmt.Scan(&numbers)
	for i = 0; i < c; i++ {
		fmt.Scan(&number)
		if (number % 8 == 0) && (number >= 10) && (number <= 99) {
			sum += number
		}
	}
	fmt.Println(sum)
}


