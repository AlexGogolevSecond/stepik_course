/*
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B  включительно.

Sample Input:

1 5
Sample Output:

15
*/
package main

import "fmt"


func main(){
	var A, B, sum int
	fmt.Scan(&A, &B)
	// fmt.Scan(&B)

	for i := A; i <= B; i++ {
		sum += i
	}
	fmt.Println(sum)
}

